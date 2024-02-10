package tests_test

import (
	"testing"
	"togglDecks/testutils"
)

func TestMain(m *testing.M) {
	testutils.TestRequiresDB(m)
}
