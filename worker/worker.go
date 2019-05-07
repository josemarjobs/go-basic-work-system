package worker

import (
	"log"
	"sync"
)

type Task interface {
	Process() error
}

type Worker struct {
	workers int
	wg      sync.WaitGroup
	jobs    chan Task
}

func NewWorker(workers int) *Worker {
	return &Worker{
		workers: workers,
		jobs:    make(chan Task),
	}
}

func (w *Worker) Start() {
	log.Printf("creating %v workers\n", w.workers)
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go func(i int) {
			log.Printf("Launching worker #%v\n", i)
			for t := range w.jobs {
				t.Process()
			}
			w.wg.Done()
			log.Printf("Shutting down worker #%v\n", i)
		}(i + 1)
	}

	log.Printf("%v workers created\n", w.workers)
}

func (w *Worker) Run(t Task) {
	w.jobs <- t
}

func (w *Worker) Shutdown() {
	close(w.jobs)
	w.wg.Wait()
}
