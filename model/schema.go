package model

type Schema struct {
	errors     []Error
	namespaces namespaces
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
