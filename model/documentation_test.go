package model

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshallDocumentationUnknownAttr(t *testing.T) {
	data := []byte(`<xs:documentation lang="eng" xml:source="test" otherAttr="test2">Washington</xs:documentation>`)
	documentation := New_Documentation(&Schema{})
	err := xml.Unmarshal(data, &documentation)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(documentation.root.errors))
}

func Test_UnmarshallDocumentationSyntaxError(t *testing.T) {
	data := []byte(`<xs:documentation>Washington</xs:documentation`)
	documentation := New_Documentation(&Schema{})
	err := xml.Unmarshal(data, &documentation)
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(documentation.root.errors))
}

func Test_UnmarshallDocumentation(t *testing.T) {
	data := []byte(`<xs:documentation xml:lang="eng" source="test"><myTag>Washington</myTag></xs:documentation>`)
	documentation := New_Documentation(&Schema{})
	err := xml.Unmarshal(data, &documentation)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(documentation.root.errors))
	assert.Equal(t, "test", documentation.source)
	assert.Equal(t, "eng", documentation.lang)
	assert.Equal(t, "Washington", documentation.content)

	documentation1 := New_Documentation(&Schema{})
	err = xml.Unmarshal([]byte(documentation.X(0)), &documentation1)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(documentation1.root.errors))
	assert.Equal(t, "test", documentation1.source)
	assert.Equal(t, "eng", documentation1.lang)
	assert.Equal(t, "Washington", documentation1.content)
}

func Test_UnmarshallDocumentationWithoutContent(t *testing.T) {
	data := []byte(`<xs:documentation source="test"/>`)
	documentation := New_Documentation(&Schema{})
	err := xml.Unmarshal(data, &documentation)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(documentation.root.errors))
	assert.Equal(t, "test", documentation.source)
	assert.Equal(t, 0, len(documentation.content))

	documentation1 := New_Documentation(&Schema{})
	err = xml.Unmarshal([]byte(documentation.X(0)), &documentation1)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(documentation1.root.errors))
	assert.Equal(t, "test", documentation1.source)
	assert.Equal(t, 0, len(documentation.content))
}
