package model

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testElem struct{
	precondition func()
	baseXsd
	unmarshalFunc func (d *xml.Decoder, start xml.StartElement) error
}

func (e *testElem) init(){
	e.initialize()
	e.root = Schema{}
	e.root.namespaces = New_Namespaces()
	e.precondition()
}

func (e *testElem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{
	e.unmarshalFunc(d, start)
	return nil
}


func Test_NameSpace(t *testing.T){
	ns:=New_Namespaces()
	err:=ns.add("prefix1", "uri1")
	assert.Nil(t, err)
	err = ns.add("prefix2", "uri2")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(ns.prefixToUri))
	assert.Equal(t, "prefix1", ns.getPrefix("uri1"))
	assert.Equal(t, "uri2", ns.getUri("prefix2"))
	err = ns.add("prefix2", "uri2")
	assert.NotNil(t, err)
}

type testElem2 struct{
	baseParticle
}
func Test_BaseParticleComponent(t *testing.T){
	elem := testElem2{}
	elem.initialize()
	assert.NotNil(t, elem.otherAttrs)
	assert.Equal(t, 1, elem.minOccurs)
	assert.Equal(t, 1, elem.maxOccurs)
}

func Test_CustomAttr(t *testing.T){
	var elem testElem
	elem.precondition = func(){
		elem.root.namespaces.add("r", "http://schemas.openxmlformats.org/officeDocument/2006/relationships")
		elem.root.namespaces.add("xsd", "http://www.w3.org/2001/XMLSchema")
	}			
	elem.init()
	
	elem.unmarshalFunc = func(d *xml.Decoder, start xml.StartElement) error {
TokenLoop:
	for {
			tok, err:=d.Token()
			assert.Nil(t, err)

			switch token:=tok.(type){
			case xml.StartElement:
				if token.Name.Local=="attribute"{
					for _, attr:=range token.Attr{
						switch attr.Name.Local{
						case "name","type":
							err:=elem.addCustomAttr(attr)
							assert.Nil(t, err)
						default:
							err:=elem.addCustomAttr(attr)
							assert.NotNil(t, err)
						}
					}
				}
			case xml.EndElement:
				if token.Name == start.Name{
					break TokenLoop
				}
			}
		}
		return nil	
	}


	data:=[]byte(`<xsd:schema targetNamespace="http://schemas.openxmlformats.org/drawingml/2006/main" elementFormDefault="qualified" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns="http://schemas.openxmlformats.org/drawingml/2006/main">
		<xs:attribute name="key-number" r:type="xs:integer" s:jojon="coba"/>
		</xsd:schema>`)
	xml.Unmarshal(data, &elem)	
}
