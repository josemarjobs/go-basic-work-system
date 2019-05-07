package main

import (
	"log"
)

type Music struct {
	Title string
}

type MusicUploaderTask struct {
	Music *Music
}

func (t *MusicUploaderTask) Process() error {
	log.Printf("Uploading %v\n", t.Music)
	return nil
}
