/*
 * @Descripttion:
 * @version:
 * @Author: WangShuaibing
 * @Date: 2020-11-30 13:46:23
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-11-30 13:54:26
 */
package code

func SquareRoot(data float64, precision float64) float64 {
	var (
		low  float64
		high float64
		mid  float64
		tmp  float64
	)

	if data > 1 {
		low = 1
		high = data
	} else {
		low = data
		high = 1
	}

	for low <= high {
		mid = (low + high) / 2.000
		tmp = mid * mid
		if tmp-data <= precision && tmp-data >= precision*-1 {
			return mid
		} else if tmp > data {
			high = mid
		} else {
			low = mid
		}
	}
	return -1.000
}
