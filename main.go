// Package amiworking can watch /etc/resolv.conf for your company domain
// and log if your working or not.
package amiworking

import (
	"bufio"
	"os"
	"regexp"
	"time"

	"gopkg.in/fsnotify.v1"
)

// Watch a file for a domain
func Watch(file, domain string, events chan bool) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer func() { _ = watcher.Close() }()

	err = watcher.Add(file)
	if err != nil {
		return err
	}

	var regex = regexp.MustCompile("(domain|search) .*(" + domain + ").*")
	result, err := check(file, regex)
	if err != nil {
		return err
	}
	events <- result
	return loop(domain, regex, watcher, events)
}

func loop(
	domain string,
	regex *regexp.Regexp,
	watcher *fsnotify.Watcher,
	events chan bool,
) error {
	for {
		var event = <-watcher.Events
		if event.Op&fsnotify.Write == fsnotify.Write {
			result, err := check(event.Name, regex)
			if err != nil {
				return err
			}
			events <- result
		} else if event.Op&fsnotify.Remove == fsnotify.Remove {
			events <- false
			for {
				time.Sleep(10 * time.Second)
				if _, err := os.Stat(event.Name); err == nil {
					return Watch(event.Name, domain, events)
				}
			}
		}
	}
}

func check(filename string, regex *regexp.Regexp) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if matches := regex.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
			return true, nil
		}
	}
	return false, nil
}
