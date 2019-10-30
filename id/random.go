package id

import (
	"math/rand"
	"time"
)

// 生成随机id
// 参考 https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

func init() {
	rand.Seed(time.Now().UnixNano())
}
