package bank

import (
	"testing"
)

// case1
func TestShareCond(t *testing.T) {
	go teller()
}

// case2
func TestShareLock(t *testing.T) {

}
