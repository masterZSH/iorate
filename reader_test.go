package iorate

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	var buf bytes.Buffer
	for i := 0; i < 1000; i++ {
		buf.WriteByte(1)
	}
	reader := NewReader(&buf, 100)
	// reader.SetLimit(100)
	start := time.Now()
	bf := make([]byte, 10)
	for {
		_, err := reader.Read(bf)
		// n, err := file.Read(buf)
		if err != nil {
			break
		}
	}
	end := time.Now()
	dur := end.Sub(start)
	assert.Equal(t, 10, int(dur.Seconds()))
}

func Test1(t *testing.T) {
	var buf bytes.Buffer
	for i := 0; i < 10*MB; i++ {
		buf.WriteByte(1)
	}
	// 1MB/s
	reader := NewReader(&buf, MB)
	start := time.Now()
	bf := make([]byte, 10)
	for {
		_, err := reader.Read(bf)
		if err != nil {
			break
		}
	}
	end := time.Now()
	dur := end.Sub(start) // 10s
	_, _ = end, dur
}
