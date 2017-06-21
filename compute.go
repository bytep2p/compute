package compute

import (
	"math"
)

// 地球半径，单位米
const R = 6367000

// lonA, latA分别为A点的纬度和经度
// lonB, latB分别为B点的纬度和经度
// 返回的距离单位为米
func Sphere(lonA, latA, lonB, latB float64) float64 {
	c := math.Sin(latA)*math.Sin(latB)*math.Cos(lonA-lonB) + math.Cos(latA)*math.Cos(latB)
	return R * math.Acos(c) * math.Pi / 180
}
