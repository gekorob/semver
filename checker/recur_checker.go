package checker

import (
	"fmt"
	"strconv"
	"strings"
)

type RecurCompare struct {
	name  string
	parts int
}

func NewRecurCompare(partsCount int) *RecurCompare {
	rCmp := RecurCompare{
		name:  "recursive compare",
		parts: partsCount,
	}

	return &rCmp
}

func (c *RecurCompare) Compare(v, w string) (n int, err error) {
	switch {
	case !c.IsValid(v):
		return -1, fmt.Errorf("Invalid version string %s", v)
	case !c.IsValid(w):
		return 1, fmt.Errorf("Invalid version string %s", w)
	}

	varr, _ := slice(v, c.parts)
	warr, _ := slice(w, c.parts)

	cmp := recurCompare(varr, warr)

	return cmp, err
}

func (c *RecurCompare) IsValid(v string) bool {
	_, err := slice(v, c.parts)
	if err != nil {
		return false
	}
	return true
}

func recurCompare(varr, warr []int64) int {
	if len(varr) == 0 {
		return 0
	}

	v := varr[0]
	w := warr[0]

	switch {
	case v > w:
		return 1
	case v < w:
		return -1
	}

	return recurCompare(varr[1:], warr[1:])
}

func slice(v string, partsCount int) (ns []int64, err error) {
	versNumbers := strings.SplitN(v, ".", partsCount)

	converted := make([]int64, partsCount, partsCount)
	for i, e := range versNumbers {
		val, err := strconv.ParseInt(e, 10, 64)
		converted[i] = val
		if err != nil {
			return nil, err
		}
	}

	return converted, nil
}
