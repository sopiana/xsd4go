package model

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHelper struct {
	unmarshaller map[string]xml.Unmarshaler
}

func New_TestHelper(s string, x xml.Unmarshaler) testHelper {
	ret := testHelper{
		unmarshaller: make(map[string]xml.Unmarshaler),
	}
	ret.unmarshaller[s] = x
	return ret
}

func (x *testHelper) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
TokenLoop:
	for {
		tok, _ := d.Token()
		switch token := tok.(type) {
		case xml.StartElement:
			obj := x.unmarshaller[token.Name.Local]
			obj.UnmarshalXML(d, token)
		case xml.EndElement:
			if token.Name == start.Name {
				break TokenLoop
			}
		}
	}
	return nil
}

func makeTestSchema(s string) []byte {
	return []byte(fmt.Sprintf(`<xs:schema targetNamespace="http://myschema.com" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns="http://myschema.com">
	%s</xs:schema>`, s))
}

func Test_getIndent(t *testing.T) {
	assert.Equal(t, "", getIndent(0))
	assert.Equal(t, "  ", getIndent(1))
	assert.Equal(t, "        ", getIndent(4))
}
