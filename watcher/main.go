package watcher

import (
	"errors"
	"log"

	"gopkg.in/fsnotify.v1"
)

func Watch(domain string) error {
	if domain == "" {
		return errors.New("Invalid domain")
	}
	log.Println("Watching for domain", domain)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add("/etc/resolv.conf"); err != nil {
		return err
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("File Modified:", event.Name)
				//regex
				//log se bate na regex
			}
		case err := <-watcher.Errors:
			log.Println(err)
		}
	}
}
