package watcher

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"

	"gopkg.in/fsnotify.v1"
)

func Watch(domain string) error {
	if domain == "" {
		return errors.New("Invalid domain")
	}
	log.Println("Watching for domain", domain)
	return watch(domain)
}

func watch(domain string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add("/etc/resolv.conf"); err != nil {
		return err
	}

	regex := regexp.MustCompile("domain .*(" + domain + ").*")
	return loop(domain, regex, watcher)
}

func loop(
	domain string,
	regex *regexp.Regexp,
	watcher *fsnotify.Watcher,
) error {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				if err := logDomain(event, regex); err != nil {
					return err
				}
			}
		case err := <-watcher.Errors:
			log.Println(err)
		}
	}
}

func logDomain(event fsnotify.Event, regex *regexp.Regexp) error {
	file, err := os.Open(event.Name)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if matches := regex.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
			log.Println("Into domain", matches[1])
		}
	}
	return nil
}
