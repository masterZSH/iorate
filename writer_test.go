package iorate

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	var buf bytes.Buffer
	// per second
	writer := NewWriter(&buf, 100)
	// reader.SetLimit(100)
	start := time.Now()
	for i := 0; i < 100; i++ {
		_, err := writer.Write([]byte("0123456789"))
		if err != nil {
			break
		}
	}
	end := time.Now()
	dur := end.Sub(start)
	assert.Equal(t, 10, int(dur.Seconds()))
}
