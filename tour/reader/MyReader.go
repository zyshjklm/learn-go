// emits an infinite stream of the ASCII character 'A'.
package main 

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct {}

func (my MyReader) Read(b []byte) (int, error) {
	if b == nil || len(b) == 0 {
		err := fmt.Errorf("Buffer is not long enough")
		return 0, err
	}
	for i, _ := range b {
		b[i] = 'A'
	}
	return len(b), nil
}


func main() {
	reader.Validate(MyReader{})
	// 
}

