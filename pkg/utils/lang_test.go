package utils

import (
	"context"
	"testing"
	"unsafe"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestOmit(t *testing.T) {
	sample := map[string]string{
		"UTC": "Universal Time Coordinated",
		"EST": "Eastern Time",
		"CET": "Central Europe Time",
	}

	// Omit single key
	func() {
		omit := Omit(sample, "UTC")
		_, utcFound := omit["UTC"]
		_, cetFound := omit["CET"]
		_, estFound := omit["EST"]
		assert.Equal(t, len(sample)-1, len(omit))
		assert.NotSame(t, &sample, &omit, "no copy made")
		assert.Equal(t, true, cetFound, "not kept")
		assert.Equal(t, true, estFound, "not kept")
		assert.Equal(t, false, utcFound, "not omitted")
	}()

	// Omit more keys
	func() {
		omit := Omit(sample, "EST", "CET")
		utc, utcFound := omit["UTC"]
		_, cetFound := omit["CET"]
		_, estFound := omit["EST"]
		assert.Equal(t, len(sample)-2, len(omit))
		assert.NotSame(t, &sample, &omit, "no copy made")
		assert.Equal(t, false, cetFound, "not omitted")
		assert.Equal(t, false, estFound, "not omitted")
		assert.Equal(t, true, utcFound, "not kept")
		assert.Equal(t, "Universal Time Coordinated", utc, "changed value")
	}()

	func() {
		omit := Omit(sample)
		_, utcFound := omit["UTC"]
		_, cetFound := omit["CET"]
		_, estFound := omit["EST"]
		assert.Equal(t, len(sample), len(omit))
		assert.NotSame(t, &sample, &omit, "no copy made")
		assert.Equal(t, true, cetFound, "not kept")
		assert.Equal(t, true, utcFound, "not kept")
		assert.Equal(t, true, estFound, "not kept")
	}()
}

func TestContainsString(t *testing.T) {
	sample := []string{"UTC", "EST", "CET"}
	assert.Equal(t, true, ContainsString(sample, "UTC"))
	assert.Equal(t, false, ContainsString(sample, " UTC "))
	assert.Equal(t, false, ContainsString(sample, ""))
}

func TestIsNumber(t *testing.T) {
	numbers := []interface{}{
		int8(0),
		int(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		uintptr(0),
		float32(0.0),
		float64(0.0),
		complex64(0),
		complex128(0),
	}

	for _, i := range numbers {
		assert.Equal(t, true, IsNumber(i))
	}

	i := 0
	ch := make(chan int)
	mMap := make(map[int]int)
	slice := make([]int, 0)
	notNumbers := []interface{}{
		bool(false),
		[]int{},
		ch,
		func() {},
		mMap,
		&i,
		slice,
		"abc",
		struct{}{},
		unsafe.Pointer(&i),
	}

	for _, nan := range notNumbers {
		assert.Equal(t, false, IsNumber(nan))
	}
}

func TestDifference(t *testing.T) {
	assert.Equal(
		t,
		[]string{"abc"},
		Difference([]string{"abc", "def", "ghi"}, []string{"def", "ghi"}),
	)
	assert.Equal(
		t,
		[]string{"abc", "def"},
		Difference([]string{"abc", "def", "ghi"}, []string{"ghi"}),
	)
	assert.Equal(
		t,
		[]string{"abc", "def", "ghi"},
		Difference([]string{"abc", "def", "ghi"}, []string{}),
	)
	assert.Equal(
		t,
		[]string{"abc", "ghi"},
		Difference([]string{"abc", "def", "ghi"}, []string{"def"}),
	)
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, -1, Min(-1, 2))
	assert.Equal(t, 0, Min(0, 1))
}

func TestDefaultContext(t *testing.T) {
	var ctx context.Context
	ctx, _ = DefaultContext()
	_, deadlineSet := ctx.Deadline()
	assert.Equal(t, true, deadlineSet, "deadline not set")
}

func TestNumberNormalizer(t *testing.T) {
	assert.Equal(t, true,
		cmp.Equal(
			[]interface{}{1, 2.0, "3", "4", 5.1},
			[]interface{}{"1", "2", 3, 4.0, "5.1"},
			NumberNormalizer))
}
