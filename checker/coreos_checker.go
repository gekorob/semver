package checker

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
)

type CoreosCompare struct {
	name string
}

func NewCoreosCompare() *CoreosCompare {
	cCmp := CoreosCompare{
		name: "coreos compare",
	}

	return &cCmp
}

func (c *CoreosCompare) Compare(v, w string) (n int, err error) {

	switch {
	case !c.IsValid(v):
		return -1, fmt.Errorf("Invalid version string %s", v)
	case !c.IsValid(w):
		return 1, fmt.Errorf("Invalid version string %s", w)
	}

	vver := semver.New(v)
	wver := semver.New(w)

	return vver.Compare(*wver), nil
}

func (c *CoreosCompare) IsValid(v string) bool {
	_, err := semver.NewVersion(v)
	if err != nil {
		return false
	}
	return true
}
