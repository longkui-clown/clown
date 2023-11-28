package utils

import (
	"fmt"
	"strconv"
)

// 8位
type GuidTag uint8

type GuidDef struct {
	Tag    GuidTag
	Retain uint8
	Id     uint
}

// 用字符串避免一些问题
func ToGuid(tag GuidTag, id uint) string {
	// 8位 渠道标识 + 8位余留 + 48位帐号id
	return fmt.Sprintf("%x", (uint(tag)<<48)|(id&(0xFFFFFFFFFF)))
}

func FromGuid(s_guid string) (*GuidDef, error) {
	guid, err := strconv.ParseInt(s_guid, 16, 64)
	if err != nil {
		return nil, err
	}
	return &GuidDef{
		Tag:    GuidTag(guid >> (48 + 8) & 0xF),
		Retain: uint8(guid >> 48 & 0xF),
		Id:     uint(guid & (0xFFFFFFFFFF))}, nil
}
