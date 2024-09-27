package utils

import "math"

// 对float 四舍五入 保留两位小数
func RoundFloat64(f, digit float64) float64 {
	factor := math.Pow(10, digit)
	return math.Round(f*factor) / factor
}
