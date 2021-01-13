package golang_demo

import (
	"fmt"
	"os"
)

func Demo() {
	scmEnv := os.Getenv("SCM_ENV")
	fmt.Println(scmEnv)
}
