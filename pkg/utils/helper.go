package utils

import (
	"net"
	"reflect"
	"strconv"
)

// 集合运算
// intersect 获取交
func Intersect[T comparable](a []T, b []T) []T {
	inter := make([]T, 0)
	mp := make(map[T]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}

// minus 获取差集
func Minus[T comparable](a []T, b []T) []T {
	var inter []T
	mp := make(map[T]bool)
	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if mp[s] {
			delete(mp, s)
		}
	}
	for key := range mp {
		inter = append(inter, key)
	}
	return inter
}

func ContainsItem[T comparable](arr []T, item T) bool {
	for _, a := range arr {
		if item == a {
			return true
		}
	}
	return false
}

// 手机号隐藏
func MaskPhone(input string) string {
	return input[:3] + "****" + input[len(input)-4:]
}

// 切片反转
func SliceReverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func SliceEqual[T any](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	return reflect.DeepEqual(s1, s2)
}

func IntToStr[T Number](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func StrToInt(i string) int64 {
	val, _ := strconv.ParseInt(i, 10, 64)
	return val
}

func StrToFloat(i string) float64 {
	val, _ := strconv.ParseFloat(i, 64)
	return val
}

func ToStringSlice[T Number](s []T) []string {
	new := []string{}
	for _, v := range s {
		new = append(new, IntToStr(v))
	}
	return new
}

// 获取本机的MAC地址
func GetMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	inter := interfaces[0]
	mac := inter.HardwareAddr.String()
	return mac
}
