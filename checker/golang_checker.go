package checker

import (
	"fmt"

	"golang.org/x/mod/semver"
)

type GolangCompare struct {
	name string
}

func NewGolangCompare() *GolangCompare {
	gCmp := GolangCompare{
		name: "golang compare",
	}

	return &gCmp
}

func (c *GolangCompare) Compare(v, w string) (n int, err error) {
	switch {
	case !c.IsValid(v):
		return -1, fmt.Errorf("Invalid version string %s", v)
	case !c.IsValid(w):
		return 1, fmt.Errorf("Invalid version string %s", w)
	}

	return semver.Compare(v, w), nil
}

func (c *GolangCompare) IsValid(v string) bool {
	return semver.IsValid(v)
}
