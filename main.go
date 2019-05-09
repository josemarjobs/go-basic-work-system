package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	worker := NewWorker(3)
	worker.Start()

	log.Println("Firing jobs")
	time.Sleep(4 * time.Second)
	for i := 1; i <= 10; i++ {
		m := &Music{Title: fmt.Sprintf("Title #%v", i)}
		worker.Run(&MusicUploaderTask{Music: m})
	}

	fmt.Print("\n\n")
	log.Println("shutting down worker system")
	time.Sleep(4 * time.Second)
	worker.Shutdown()
}
