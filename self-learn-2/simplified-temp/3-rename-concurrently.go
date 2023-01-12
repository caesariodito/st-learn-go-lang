package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

type FileInfo struct {
	Index       int
	FilePath    string
	FileName    string
	Content     []byte
	WorkerIndex int
	Err         error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start")
	start := time.Now()

	run()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func run() {
	// pipeline 1: job distribution
	chanFileIndex := readFiles()

	// for n := range chanFileIndex {
	// 	fmt.Printf("n = %s\n", n.FilePath)
	// }

	// pipeline 2: the main logic (creating files)
	createFilesWorker := 100
	chanFileResult := renameAndCalculate(chanFileIndex, createFilesWorker)

	// track and print output
	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}

	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		i := 0
		filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
			// if it is a sub directory, return immediatelly
			if info.IsDir() {
				return nil
			}

			buf, err := ioutil.ReadFile(path)

			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
				Index:    i,
				Err:      err,
			}

			i++

			return nil
		})

		close(chanOut)
	}()

	return chanOut
}

func renameAndCalculate(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// distribute job
	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorkers)

	go func() {
		// dispatching
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {

				for fileInfo := range chanIn {
					// log
					// log.Println("worker", workerIndex, "working on", fileInfo.FilePath, "file calculate and renaming")

					// calculate
					fileInfo.FileName = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))

					// rename
					newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.FileName))
					err := os.Rename(fileInfo.FilePath, newPath)
					fileInfo.FilePath = newPath
					fileInfo.Err = err
					fileInfo.WorkerIndex = workerIndex

					chanOut <- fileInfo
				}
				wg.Done()

			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
