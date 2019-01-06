package looker

import (
	"reflect"
)

type Package struct {
	Parameters []StructElement
}

func LookAtParameters(prms []reflect.Type) Package {
	pkg := Package{
		Parameters: make([]StructElement, 0, len(prms)),
	}
	for _, typ := range prms {
		prm := LookAtParameter(typ)
		pkg.Parameters = append(pkg.Parameters, prm)
	}

	return pkg
}

type StructElement struct {
	UserType string
	Fields   Fields
}

func LookAtParameter(at reflect.Type) StructElement {
	if at.Kind() == reflect.Ptr {
		at = at.Elem()
	}
	var prm = StructElement{
		UserType: at.Name(),
		Fields:   LookAtFields(at),
	}

	return prm
}

const (
	TagQuery  = "query"
	TagHeader = "header"
	TagJson   = "json"
	TagPath   = "path"
)

var tags = [...]string{
	TagQuery,
	TagHeader,
	TagPath,
	TagJson,
}

type Field struct {
	Name     string
	BaseType string
	TagName  string
	TagValue string
}

type Fields []Field

func LookAtFields(st reflect.Type) Fields {
	fields := make(Fields, 0, st.NumField())
	for i := 0; i < st.NumField(); i++ {
		ft := st.Field(i)
		fields = append(fields, LookAtField(ft)...)
	}
	return fields
}

func LookAtField(ft reflect.StructField) Fields {
	f := Field{
		Name:     ft.Name,
		BaseType: ft.Type.Kind().String(),
	}
	for _, tagName := range tags {
		tagValue, ok := ft.Tag.Lookup(tagName)
		if ok {
			f.TagName = tagName
			f.TagValue = tagValue
			break
		}
	}
	return []Field{f}
}
