## iorate

### Reader

```go
package yours

import (
  "github.com/masterZSH/iorate"
)

func FooFunc() {
    var buf bytes.Buffer
	for i := 0; i < 10*iorate.MB; i++ {
		buf.WriteByte(1)
	}
    // 1MB/s
	reader := iorate.NewReader(&buf, iorate.MB)
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
        _, _ = end,dur
}
```


### Writer

```go
package yours

import (
  "github.com/masterZSH/iorate"
)

func FooFunc() {
    var buf bytes.Buffer
	// per second 1KB
	writer := iorate.NewWriter(&buf, iorate.KB)
	start := time.Now()
	for i := 0; i < 10*iorate.KB; i++ {
		_, err := writer.Write([]byte("1"))
		if err != nil {
			break
		}
	}
	end := time.Now()
	dur := end.Sub(start) // 10s
        _, _ = end,dur
}
```
