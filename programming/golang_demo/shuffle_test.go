package golang_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestShuffleInts(t *testing.T) {
	var tmpSlice = []int{0, 1, 2, 4, 5, 6, 7, 8}
	ShuffleInts(tmpSlice)
	fmt.Printf("kingScan %+v", tmpSlice)

	fmt.Printf("kingScan %+v", time.Now().UTC().UnixNano())
}
