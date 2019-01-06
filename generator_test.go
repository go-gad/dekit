package main

import (
	"io/ioutil"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var update bool = false

func TestGenerateCode(t *testing.T) {
	code, err := GenerateCode("github.com/go-gad/dekit/looker/testdata", []string{"CreateOrderReq"})
	if err != nil {
		t.Fatalf("Failed to generate a code: %+v", err)
	}

	t.Logf("Generator output:\n%s\n", code)

	if update {
		if err = ioutil.WriteFile("./looker/testdata/decoders.go", code, 0666); err != nil {
			t.Fatalf("failed to write file: %+v", err)
		}
	}

	expCode, err := ioutil.ReadFile("./looker/testdata/decoders.go")
	if string(expCode) != string(code) {
		t.Error("generated code is not equal to expected")
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(string(expCode), string(code), true)
		t.Log(dmp.DiffPrettyText(diffs))
	}
}
