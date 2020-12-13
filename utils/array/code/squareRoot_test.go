/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-11-30 13:46:23
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-11-30 14:02:00
 */
package code

import (
	"fmt"
	"math"
	"testing"
)

func TestSquareRoot(t *testing.T) {
	num := SquareRoot(0.09, 0.000001)
	fmt.Printf("%f \n", num)

	var NewNum = 0.09
	sqrtNum := math.Sqrt(float64(NewNum))
	fmt.Printf("%f \n", sqrtNum)

}
