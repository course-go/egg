package egg_test

import (
	"testing"

	"github.com/course-go/egg/internal/egg"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"egg": egg.Run,
	})
}

func TestEgg(t *testing.T) {
	t.Parallel()

	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
