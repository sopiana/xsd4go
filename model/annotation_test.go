package model

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshallAnnotationInvalidCharData(t *testing.T) {
	data := []byte(`<xs:schema targetNamespace="http://myschema.com" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns="http://myschema.com">
					<xs:annotation id="7826" name="MyName">
						<!--this is just comment-->
						<xs:documentation>States in the Pacific Northwest of US</xs:documentation>
						IllegalChar
					</xs:annotation>
		</xs:schema>`)
	s := New_Schema("test.xsd")
	annotation := New_Annotation(&s, nil)
	helper := New_TestHelper("annotation", &annotation)
	err := xml.Unmarshal(data, &helper)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(s.errors))
	for _, e := range s.errors {
		fmt.Println(fmt.Sprint(e))
	}
}

func Test_UnmarshallAnnotationInvalidChildElem(t *testing.T) {
	data := []byte(`<xs:schema targetNamespace="http://myschema.com" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns="http://myschema.com">
					<xs:annotation id="7826" name="MyName">
						<xs:documentation>States in the Pacific Northwest of US</xs:documentation>
						<xs:all/>
					</xs:annotation>
		</xs:schema>`)
	s := New_Schema("test.xsd")
	annotation := New_Annotation(&s, nil)
	helper := New_TestHelper("annotation", &annotation)
	err := xml.Unmarshal(data, &helper)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(s.errors))
	for _, e := range s.errors {
		fmt.Println(fmt.Sprint(e))
	}
}

func Test_UnmarshallAnnotationInvalidAttrNs(t *testing.T) {
	data := []byte(`<xs:schema targetNamespace="http://myschema.com" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns="http://myschema.com">
					<xs:annotation id="7826" cd:name="MyName">
						<xs:documentation>States in the Pacific Northwest of US</xs:documentation>
					</xs:annotation>
		</xs:schema>`)
	s := New_Schema("test.xsd")
	annotation := New_Annotation(&s, nil)
	helper := New_TestHelper("annotation", &annotation)
	err := xml.Unmarshal(data, &helper)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(s.errors))
	for _, e := range s.errors {
		fmt.Println(fmt.Sprint(e))
	}
}

func Test_UnmarshallAnnotationOK(t *testing.T) {
	data := []byte(`<xs:schema targetNamespace="http://myschema.com" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns="http://myschema.com">
					<xs:annotation id="7826" name="MyName">
						<xs:documentation>States in the Pacific Northwest of US</xs:documentation>
						<xs:appinfo>MyAppInfo</xs:appinfo>
					</xs:annotation>
		</xs:schema>`)
	s := New_Schema("test.xsd")
	annotation := New_Annotation(&s, nil)
	helper := New_TestHelper("annotation", &annotation)
	err := xml.Unmarshal(data, &helper)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(s.errors))
	assert.Equal(t, "7826", annotation.id)
	assert.Equal(t, 1, len(annotation.otherAttrs))
	assert.Equal(t, "MyName", annotation.otherAttrs["name"])
	assert.Equal(t, 2, len(annotation.content))
	assert.Equal(t, "States in the Pacific Northwest of US", annotation.content[0].(Documentation).content)
	assert.Equal(t, &s, annotation.content[0].(Documentation).root)
	assert.Equal(t, "MyAppInfo", annotation.content[1].(AppInfo).content)
	assert.Equal(t, &s, annotation.content[1].(AppInfo).root)

	annotation2 := New_Annotation(&s, nil)
	_ = annotation2
	fmt.Println(annotation.X(0))
	helper = New_TestHelper("annotation", &annotation2)
	data = makeTestSchema(annotation.X(0))
	err = xml.Unmarshal(data, &helper)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(annotation, annotation2))
}

func Test_UnmarshallAnnotationEmpty(t *testing.T) {
	data := makeTestSchema(`<xs:annotation/>`)
	s := New_Schema("test.xsd")
	annotation := New_Annotation(&s, nil)
	helper := New_TestHelper("annotation", &annotation)
	err := xml.Unmarshal(data, &helper)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(s.errors))

	annotation2 := New_Annotation(&s, nil)
	_ = annotation2
	fmt.Println(annotation.X(0))
	helper = New_TestHelper("annotation", &annotation2)
	data = makeTestSchema(annotation.X(0))
	err = xml.Unmarshal(data, &helper)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(annotation, annotation2))
}
