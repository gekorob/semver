package checker_test

import (
	"testing"

	"github.com/gekorob/semver/checker"
)

func TestGolangChecker(t *testing.T) {
	gChecker := checker.NewGolangCompare()
	cmp, err := checker.Compare(gChecker, "v4.3.1", "v4.4.4")

	if err != nil {
		t.Errorf("Error checking versions %s", err)
	}

	if cmp != -1 {
		t.Errorf("Was expecting -1, got: %v", cmp)
	}
}

func TestCoreosChecker(t *testing.T) {
	cChecker := checker.NewCoreosCompare()
	cmp, err := checker.Compare(cChecker, "4.3.1", "4.4.4")

	if err != nil {
		t.Errorf("Error checking versions %s", err)
	}

	if cmp != -1 {
		t.Errorf("Was expecting -1, got: %v", cmp)
	}
}

func TestRecurChecker(t *testing.T) {
	rChecker := checker.NewRecurCompare(4)
	cmp, err := checker.Compare(rChecker, "4.4.4.4", "4.4.3")

	if err != nil {
		t.Errorf("Error checking versions %s", err)
	}

	if cmp != -1 {
		t.Errorf("Was expecting -1, got: %v", cmp)
	}
}
