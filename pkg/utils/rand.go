package utils

import (
	cr "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

var randReal bool = false
var Rander = rand.New(rand.NewSource(time.Now().UnixNano()))

const letterString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numLetterString = "0123456789"

func GetRandReal() bool {
	return randReal
}

func SetRandReal(real bool) {
	randReal = real
}

// 随机生成字符串
func RandStr(n int, letter string) string {
	str := []byte(letter)
	res := ""
	for i := 0; i < n; i++ {
		res += fmt.Sprintf("%c", str[RandNum(strings.Count(letter, "")-1)])
	}
	return res
}

func RandNumStr(n int) string {
	return RandStr(n, numLetterString)
}

func RandString(n int) string {
	return RandStr(n, letterString)
}

func RandOrder(n int) string {
	return time.Now().Format("20060102150405") + RandNumStr(n)
}

// --------------- 》》》假随机《《《 ---------------
// --------------- 》》》假随机《《《 ---------------
// 包含min, max
func FakeRandRange(min, max int) int {
	return Rander.Intn(max-min+1) + min
}

// 包含max
func FakeRandNum(max int) int {
	return FakeRandRange(0, max)
}

// 包含min, max
func FakeRandRange64(min, max int64) int64 {
	return Rander.Int63n(max-min+1) + min
}

// 包含max
func FakeRandNum64(max int64) int64 {
	return FakeRandRange64(0, max)
}

// --------------- 》》》真随机《《《 ---------------
// --------------- 》》》真随机《《《 ---------------
// 包含min, max
func RealRandRange(min, max int) int {
	n, _ := cr.Int(cr.Reader, big.NewInt(int64(max-min+1)))
	return int(n.Int64()) + min
}

// 包含max
func RealRandNum(max int) int {
	return RealRandRange(0, max)
}

// 包含min, max
func RealRandRange64(min, max int64) int64 {
	n, _ := cr.Int(cr.Reader, big.NewInt(max-min+1))
	return n.Int64() + min
}

// 包含max
func RealRandNum64(max int64) int64 {
	return RealRandRange64(0, max)
}

// -------------------------- 统一接口调用 --------------------------------
// -------------------------- 统一接口调用 --------------------------------
func RandRange(min, max int) int {
	if randReal {
		return RealRandRange(min, max)
	}
	return FakeRandRange(min, max)
}

func RandNum(max int) int {
	if randReal {
		return RealRandNum(max)
	}
	return FakeRandNum(max)
}

func RandRange64(min, max int64) int64 {
	if randReal {
		return RealRandRange64(min, max)
	}
	return FakeRandRange64(min, max)
}

func RandNum64(max int64) int64 {
	if randReal {
		return RealRandNum64(max)
	}
	return FakeRandNum64(max)
}
