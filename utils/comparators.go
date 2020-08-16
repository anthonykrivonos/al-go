package utils

import (
	"strings"
)

// AscComp is an ascending comparator for integers.
func IntAscComp(a interface{}, b interface{}) int {
	return a.(int) - b.(int)
}

// IntDescComp is an descending comparator for integers.
func IntDescComp(a interface{}, b interface{}) int {
	return IntDescComp(b, a)
}

// Float32AscComp is an ascending comparator for float32.
func Float32AscComp(a interface{}, b interface{}) int {
	if a.(float32) < b.(float32) {
		return -1
	} else if a.(float32) > b.(float32) {
		return 1
	}
	return 0
}

// Float32DescComp is an descending comparator for float32.
func Float32DescComp(a interface{}, b interface{}) int {
	return Float32DescComp(b, a)
}

// Float64AscComp is an ascending comparator for float64.
func Float64AscComp(a interface{}, b interface{}) int {
	if a.(float64) < b.(float64) {
		return -1
	} else if a.(float64) > b.(float64) {
		return 1
	}
	return 0
}

// Float64DescComp is an descending comparator for float64.
func Float64DescComp(a interface{}, b interface{}) int {
	return Float64AscComp(b, a)
}

// StringAscComp is an ascending comparator for strings.
func StringAscComp(a interface{}, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}

// StringDescComp is an descending comparator for strings.
func StringDescComp(a interface{}, b interface{}) int {
	return StringAscComp(b, a)
}
