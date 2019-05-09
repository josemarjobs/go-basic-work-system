package main

import (
	"log"
	"time"
)

type Music struct {
	Title string
}

type MusicUploaderTask struct {
	Music *Music
}

func (t *MusicUploaderTask) Process(workerId int) error {
	log.Printf("[#%v - STARTED] Uploading %v\n", workerId, t.Music)
	time.Sleep(2 * time.Second)
	log.Printf("[#%v - DONE] Uploaded %v\n", workerId, t.Music)

	return nil
}
