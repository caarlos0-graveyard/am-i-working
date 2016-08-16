package watcher_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/caarlos0/am-i-working/watcher"
	"github.com/stretchr/testify/assert"
)

func TestWatchForDomain(t *testing.T) {
	// warm-up
	assert := assert.New(t)
	file, err := ioutil.TempFile(os.TempDir(), "fake_resolv.conf")
	assert.NoError(err)
	defer os.Remove(file.Name())
	events := make(chan bool)

	// watch the file
	go watcher.Watch(file.Name(), "mydomain", events)

	// showtime
	assert.False(<-events, "the file is empty")

	assert.NoError(ioutil.WriteFile(file.Name(), []byte("domain whatever"), 0644))
	assert.False(<-events, "domain doesnt match")

	assert.NoError(ioutil.WriteFile(file.Name(), []byte("domain mydomain"), 0644))
	assert.True(<-events, "domain match")

	assert.NoError(ioutil.WriteFile(file.Name(), []byte("domain nope"), 0644))
	assert.False(<-events, "domain doesnt match")
}

func TestNonExistentFile(t *testing.T) {
	assert := assert.New(t)
	file := "/tmp/wtf-this-shouldn-exist"
	assert.Error(watcher.Watch(file, "mydomain", make(chan bool)))
}
