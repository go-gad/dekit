package looker

import (
	"testing"

	"github.com/kr/pretty"
)

func TestReflect(t *testing.T) {
	pkg, err := Reflect("github.com/go-gad/dekit/looker/testdata", []string{"CreateOrderReq"})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Package %# v", pretty.Formatter(pkg))
}
