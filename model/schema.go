package model

import "fmt"

type Schema struct {
	filename   string
	ids        map[string]Any
	errors     []Error
	namespaces namespaces
}

func New_Schema(filename string) Schema {
	ret := Schema{
		filename:   filename,
		ids:        make(map[string]Any),
		namespaces: New_Namespaces(),
	}
	ret.namespaces.add("xml", "http://www.w3.org/XML/1998/namespace")
	return ret
}

func (x *Schema) addId(id string, e Any) error {
	if _, ok := x.ids[id]; ok {
		return fmt.Errorf("id: %s is duplicated", id)
	}
	x.ids[id] = e
	return nil
}

func (x *Schema) Error(offset int64, message string) {
	x.errors = append(x.errors, Error{
		offset:  offset,
		message: message,
	})
}

type Error struct {
	offset  int64
	message string
}
