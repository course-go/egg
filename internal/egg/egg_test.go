package egg_test

import (
	"os"
	"testing"

	"github.com/course-go/egg/internal/egg"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"egg": egg.Run,
	}))
}

func TestEgg(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
