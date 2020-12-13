/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-11-10 19:49:45
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-11-10 19:53:34
 */
package concurrent_demo

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExecDemo(t *testing.T) {
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")

	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
		return
	}
}
