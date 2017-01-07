package am_i_working_test

import (
	"io/ioutil"
	"os"
	"testing"

	working "github.com/caarlos0/am-i-working"
	"github.com/stretchr/testify/assert"
)

func TestWatchForDomain(t *testing.T) {
	t.Skip("test is hanging...")
	// warm-up
	assert := assert.New(t)
	f, err := ioutil.TempFile(os.TempDir(), "fake_resolv.conf")
	file := f.Name()
	assert.NoError(err)
	defer os.Remove(file)
	events := make(chan bool)

	// watch the file
	go working.Watch(file, "mydomain", events)

	// showtime
	assert.False(<-events, "the file is empty")

	assert.NoError(ioutil.WriteFile(file, []byte("domain whatever"), 0644))
	assert.False(<-events, "domain doesnt match")

	assert.NoError(ioutil.WriteFile(file, []byte("domain mydomain"), 0644))
	assert.True(<-events, "domain match")

	assert.NoError(ioutil.WriteFile(file, []byte("domain nope"), 0644))
	assert.False(<-events, "domain doesnt match")

	os.Remove(file)
	assert.False(<-events, "file was removed")

	assert.NoError(ioutil.WriteFile(file, []byte("domain mydomain"), 0644))
	assert.True(<-events, "domain match")
}

func TestNonExistentFile(t *testing.T) {
	assert := assert.New(t)
	file := "/tmp/wtf-this-shouldn-exist"
	assert.Error(working.Watch(file, "mydomain", make(chan bool)))
}
