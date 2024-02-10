package decks

import (
	"testing"
	"togglDecks/routers"
	"togglDecks/testutils"
)

var Router = routers.SetupMainRouter()

func TestMain(m *testing.M) {
	testutils.TestRequiresDB(m)
	m.Run()
}
