package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Если значение не нулевое, вернуть максимальное значение для типа
func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {
	var (
		valueInt8  int8
		valueInt16 int16
		valueInt32 int32
		valueInt64 int64
	)
	if in8 != 0 {
		valueInt8 = ((1 << 8) - 1) / 2
	}
	if in16 != 0 {
		valueInt16 = ((1 << 16) - 1) / 2
	}
	if in32 != 0 {
		valueInt32 = ((1 << 32) - 1) / 2
	}
	if in64 != 0 {
		valueInt64 = ((1 << 64) - 1) / 2
	}

	return valueInt8, valueInt16, valueInt32, valueInt64
}

// Если значение не нулевое, вернуть максимальное значение для типа
func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	var (
		valueUint8  uint8
		valueUint16 uint16
		valueUint32 uint32
		valueUint64 uint64
	)
	if uin8 != 0 {
		valueUint8 = ((1 << 8) - 1)
	}
	if uin16 != 0 {
		valueUint16 = ((1 << 16) - 1)
	}
	if uin32 != 0 {
		valueUint32 = ((1 << 32) - 1)
	}
	if uin64 != 0 {
		valueUint64 = ((1 << 64) - 1)
	}

	return valueUint8, valueUint16, valueUint32, valueUint64

}

func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)

	return bits
}

func main() {
	// var (
	// 	in8  int8 =1
	// 	in16 int16 = 1
	// 	in32 int32 = 1
	// 	in64 int64 = 1
	// )

	// fmt.Println(getIntMaxValue(in8, in16, in32, in64))

	// var (
	// 	uin8  uint8 =1
	// 	uin16 uint16 = 1
	// 	uin32 uint32 = 1
	// 	uin64 uint64 = 1
	// )

	// fmt.Println(getUintMaxValue(uin8, uin16, uin32, uin64))

	var a int
	fmt.Println(getBits(a))
}
