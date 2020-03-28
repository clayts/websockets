package websockets

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

func TestGob(t *testing.T) {
	f := func(e interface{}) {
		buf := bytes.NewBuffer(nil)
		enc := gob.NewEncoder(buf)
		enc.Encode(e)
		fmt.Println(buf.String())
	}
	f(struct {
		s string
		i int
	}{"hi", 7})
}
