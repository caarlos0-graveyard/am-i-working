package watcher

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"gopkg.in/fsnotify.v1"
)

const resolvconf = "/etc/resolv.conf"

func Watch(domain string) error {
	log.Println("Watching for domain", domain)
	return watch(domain)
}

func watch(domain string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(resolvconf); err != nil {
		return err
	}

	regex := regexp.MustCompile("domain .*(" + domain + ").*")
	logDomain(regex)
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
				if err := logDomain(regex); err != nil {
					return err
				}
			}
		case err := <-watcher.Errors:
			log.Println(err)
		}
	}
}

func logDomain(regex *regexp.Regexp) error {
	file, err := os.Open(resolvconf)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var match string
	for scanner.Scan() {
		if matches := regex.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
			match = matches[1]
		}
	}
	if match == "" {
		log.Println("Not working")
	} else {
		log.Println("Working in domain", match)
	}
	return nil
}
