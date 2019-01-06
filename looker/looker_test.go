package looker

import (
	"reflect"
	"testing"

	"github.com/go-gad/dekit/looker/testdata"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestLookAtParameters(t *testing.T) {
	prms := []reflect.Type{
		reflect.TypeOf(testdata.Req1{}),
		reflect.TypeOf(testdata.Req2{}),
	}
	pkg := LookAtParameters(prms)
	assert.Equal(t, 2, len(pkg.Parameters))
}

func TestLookAtParameter(t *testing.T) {
	var typ reflect.Type = reflect.TypeOf(testdata.CreateOrderReq{})
	prm := LookAtParameter(typ)
	assert.Equal(t, "CreateOrderReq", prm.UserType)
	assert.Equal(t, 2, len(prm.Fields))
}

func TestLookAtFields(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		var typ reflect.Type = reflect.TypeOf(testdata.CreateOrderReq{})
		actFields := LookAtFields(typ)
		expFields := Fields{
			{
				Name:     "Name",
				BaseType: "string",
				TagName:  TagQuery,
				TagValue: "name",
			},
			{
				Name:     "Body",
				BaseType: "string",
				TagName:  TagJson,
				TagValue: "body",
			},
		}
		assert.Equal(t, expFields, actFields)
		t.Logf("struct field %# v", pretty.Formatter(actFields))
	})
}
