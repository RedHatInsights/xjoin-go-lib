package utils

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"sort"
	"time"
)

// Omit returns a copy of the given map with the given keys left out
func Omit(value map[string]string, ignoredKeys ...string) map[string]string {
	copy := make(map[string]string, len(value))

	for k, v := range value {
		ignored := false

		for _, ignoredKey := range ignoredKeys {
			if k == ignoredKey {
				ignored = true
			}
		}

		if !ignored {
			copy[k] = v
		}

	}

	return copy
}

func ContainsString(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func IsNumber(x interface{}) bool {
	kind := reflect.TypeOf(x).Kind()
	return kind >= 2 && kind <= 16
}

type void struct{}

func Difference(a, b []string) (diff []string) {
	bMap := make(map[string]void, len(b))
	diff = []string{}

	for _, key := range b {
		bMap[key] = void{}
	}

	// find missing values in a
	for _, key := range a {
		if _, ok := bMap[key]; !ok {
			diff = append(diff, key)
		}
	}

	return diff
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func SortMap(unsortedMap map[string]interface{}) map[string]interface{} {
	keys := make([]string, 0, len(unsortedMap))
	for k := range unsortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedMap := make(map[string]interface{})
	for _, k := range keys {
		sortedMap[k] = unsortedMap[k]
	}

	return sortedMap
}

func TimeFormat() string {
	return "2006-01-02T15:04:05.999999999"
}

func DefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*30)
}

// NumberNormalizer normalizes different number types (e.g. float64 vs int64) by converting them to their string representation
var NumberNormalizer = cmp.FilterValues(func(x, y interface{}) bool {
	return IsNumber(x) || IsNumber(y)
}, cmp.Transformer("NormalizeNumbers", func(in interface{}) string {
	return fmt.Sprintf("%v", in)
}))
