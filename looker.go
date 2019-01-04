package main

import (
	"reflect"
)

const (
	TagQuery  = "query"
	TagHeader = "header"
)

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
	tagValue, ok := ft.Tag.Lookup(TagQuery)
	if ok {
		f.TagName = TagQuery
		f.TagValue = tagValue
	}
	return []Field{f}
}
