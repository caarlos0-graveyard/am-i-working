package amiworking

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWatchForDomain(t *testing.T) {
	// warm-up
	var assert = assert.New(t)
	f, err := ioutil.TempFile(os.TempDir(), "fake_resolv.conf")
	assert.NoError(err)
	file := f.Name()
	defer func() { _ = os.Remove(file) }()
	events := make(chan bool)

	// watch the file
	go func() {
		assert.NoError(Watch(file, "mydomain", events))
	}()

	// showtime
	assert.False(<-events, "the file is empty")

	assert.NoError(ioutil.WriteFile(file, []byte("domain whatever"), 0644))
	assert.False(<-events, "domain doesnt match")

	assert.NoError(ioutil.WriteFile(file, []byte("domain mydomain"), 0644))
	assert.True(<-events, "domain match")

	assert.NoError(ioutil.WriteFile(file, []byte("search dev.mydomain.com"), 0644))
	assert.True(<-events, "domain match")

	assert.NoError(ioutil.WriteFile(file, []byte("domain nope"), 0644))
	assert.False(<-events, "domain doesnt match")

	assert.NoError(os.Remove(file))
	assert.False(<-events, "file was removed")

	assert.NoError(ioutil.WriteFile(file, []byte("domain mydomain"), 0644))
	assert.True(<-events, "domain match")
}

func TestNonExistentFile(t *testing.T) {
	assert := assert.New(t)
	file := "/tmp/wtf-this-shouldn-exist"
	assert.Error(Watch(file, "mydomain", make(chan bool)))
}
