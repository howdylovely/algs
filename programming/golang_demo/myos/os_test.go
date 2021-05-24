package golang_demo

import (
	"fmt"
	"os"
	"testing"
)

func TestOS(t *testing.T) {
	fmt.Println(os.Getwd())
}
