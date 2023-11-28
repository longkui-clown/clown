package utils

type Number interface {
	IntNumber | FloatNumber
}
type IntNumber interface {
	uint | int | int16 | int32 | int64 | int8
}
type FloatNumber interface {
	float32 | float64
}
