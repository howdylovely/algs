package golang_demo

import (
	"fmt"
	"testing"
)

func TestKpi(t *testing.T) {
	const JD float64 = 0.67
	const TECH float64 = 0.33

	var worktime float64
	worktime = 8

	fmt.Printf("JD: %f \n", worktime*JD)
	fmt.Printf("TECH: %f \n", worktime*TECH)
}
