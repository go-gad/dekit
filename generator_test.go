package main

import "testing"

func TestGenerateCode(t *testing.T) {
	code, err := GenerateCode("github.com/go-gad/dekit/looker/testdata", []string{"CreateOrderReq"})
	if err != nil {
		t.Fatalf("Failed to generate a code: %+v", err)
	}

	t.Logf("Generator output:\n%s\n", code)
}
