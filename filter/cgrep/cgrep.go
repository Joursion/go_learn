package cgrep

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"bytes"
)

/*page 260*/

type Result struct {
	filename string
	lino     int
	line     string
}

type Job struct {
	filename string
	results  chan<- Result
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <regexp> <files>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		grep(lineRx, commdLineFiles(os.Args[2:]))
	}
}

var workers = runtime.NumCPU()

func grep(lineRx *regexp.Regexp, filename []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(filename)))

	done := make(chan struct{}, workers)

	go addJobs(jobs, filename, results)
	for i := 0; i < workers; i++ {
		go doJobs(done, lineRx, jobs)
	}

	go awaitCompletion(done, results)
	processResults(results)
}

func addJobs(jobs chan<- Job, filenames []string, result chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job(filename, results)
	}
	close(jobs)
}

func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func NewBuf () {
	a := bytes.NewBufferString("1")
}