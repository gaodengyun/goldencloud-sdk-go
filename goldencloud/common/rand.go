package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	numLen = 10 //随机序列长度
)

var (
	numeric = [numLen]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //随机序列
)

func GetRand(width int) string {
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
