package main

import (
	"reflect"
	"testing"

	"github.com/go-gad/dekit/testdata"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestLookAtFields(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		var typ reflect.Type = reflect.TypeOf(testdata.CreateOrderReq{})
		actFields := LookAtFields(typ)
		expFields := Fields{
			{
				Name:     "Name",
				BaseType: "string",
			},
			{
				Name:     "Body",
				BaseType: "string",
			},
		}
		assert.Equal(t, expFields, actFields)
		t.Logf("struct field %# v", pretty.Formatter(actFields))
	})
}
