package am_i_working

import (
	"bufio"
	"os"
	"regexp"

	"gopkg.in/fsnotify.v1"
)

func Watch(file, domain string, events chan bool) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(file); err != nil {
		return err
	}

	regex := regexp.MustCompile("domain .*(" + domain + ").*")
	result, err := check(file, regex)
	if err != nil {
		return err
	}
	events <- result
	return loop(domain, regex, watcher, events)
}

func loop(domain string, regex *regexp.Regexp, watcher *fsnotify.Watcher, events chan bool) error {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				result, err := check(event.Name, regex)
				if err != nil {
					return err
				}
				events <- result
			}
		}
	}
}

func check(filename string, regex *regexp.Regexp) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if matches := regex.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
			return true, nil
		}
	}
	return false, nil
}
