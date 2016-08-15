package watcher

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"gopkg.in/fsnotify.v1"
)

func Watch(file, domain string) error {
	log.Println("Watching for domain", domain)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(file); err != nil {
		return err
	}

	regex := regexp.MustCompile("domain .*(" + domain + ").*")
	logDomain(file, regex)
	return loop(domain, regex, watcher)
}

func loop(domain string, regex *regexp.Regexp, watcher *fsnotify.Watcher) error {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				if err := logDomain(event.Name, regex); err != nil {
					return err
				}
			}
		case err := <-watcher.Errors:
			log.Println(err)
		}
	}
}

func logDomain(name string, regex *regexp.Regexp) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var match string
	for scanner.Scan() {
		if matches := regex.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
			match = matches[1]
			break
		}
	}
	if match == "" {
		log.Println("Not working")
	} else {
		log.Println("Working in domain", match)
	}
	return nil
}
