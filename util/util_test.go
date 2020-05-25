package util

import (
	"reflect"
	"testing"
)

func TestBinaryToIntArray(t *testing.T) {
	cases := []struct {
		bin      string
		expected []int
	}{
		{bin: "00000", expected: []int{0, 0, 0, 0, 0}},
		{bin: "11111", expected: []int{1, 1, 1, 1, 1}},
		{bin: "zabcd", expected: []int{1, 1, 1, 1, 1}},
		{bin: "00101", expected: []int{1, 0, 1, 0, 0}},
		{bin: "10000", expected: []int{0, 0, 0, 0, 1}},
	}

	for _, c := range cases {
		actual := BinaryToIntArray(c.bin)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Logf("BinaryToIntArray(%s) expected: %v got: %v", c.bin, c.expected, actual)
			t.Fail()
		}
	}
}

func TestIntArrayToBinary(t *testing.T) {
	cases := []struct {
		vals     []int
		expected string
	}{
		{expected: "00000", vals: []int{0, 0, 0, 0, 0}},
		{expected: "11111", vals: []int{1, 2, 3, 4, 5}},
		{expected: "11111", vals: []int{1, 1, 1, 1, 1}},
		{expected: "00101", vals: []int{1, 0, 1, 0, 0}},
		{expected: "10000", vals: []int{0, 0, 0, 0, 1}},
	}

	for _, c := range cases {
		actual := IntArrayToBinary(c.vals)
		if actual != c.expected {
			t.Logf("IntArrayToBinary(%v) expected: %s got: %s", c.vals, c.expected, actual)
			t.Fail()
		}
	}
}
