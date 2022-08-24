package model

import (
	"encoding/xml"
	"fmt"
)

//interface fo all xsd elements
type Any interface{
	//get inner xml file
	X(indent uint8) string
	//get go code
	Go(indent uint8) string
}

//container for non-schema attributes
type CustomAttributes map[string]string

//get inner string for otherAttrs
func (c CustomAttributes)X()string{
	var ret string
	for k, v :=range c{
		ret = fmt.Sprintf(` %s="%s"`, k, v)
	} 
	return ret
}

//add attribute into otherAttrs field. This function will return error if the namespace is not defined in xsd file 
func (c CustomAttributes)add(attr xml.Attr, ns namespaces) error{
	if len(attr.Name.Space)==0{
		c[attr.Name.Local] = attr.Value
		return nil
	}

	if att:=ns.getPrefix(attr.Name.Space); len(att)==0{
		return fmt.Errorf("invalid namespace")
	}else{
		c[att] = attr.Value
	}
	return nil
}

//Base type for all xsd component, except schema and annotation element
type baseXsd struct{
	//Root Schema Element
	root Schema

	//Parent Node
	parent Any

	//Optional: The ID of this element. The id value must be of type ID and be unique within the document containing this element.
	id string

	//{any attributes with non-schema Namespace...}>	
	otherAttrs CustomAttributes

	annotation *Annotation
}

//initialize the otherAttrs
func (b *baseXsd) initialize(){
	b.otherAttrs = make(CustomAttributes)
}

//Add custom attributes from xsd
func (b baseXsd)addCustomAttr(attr xml.Attr) error{
	return b.otherAttrs.add(attr, b.root.namespaces)
}

//Base type for Identity Constraints elements: field, key, keyref, selector, unique and element
type baseParticle struct{
	baseXsd
	
	//Optional The minimum number of times the any element can occur on the element. The value can be an integer greater than or equal to zero. 
	//To specify that this any group is optional, set this attribute to zero. Default value is 1.
	minOccurs int
	
	//Optional The maximum number of times the any element can occur on the element. 
	//The value can be an integer greater than or equal to zero. 
	//To set no limit on the maximum number, use the string "unbounded". Default value is 1.
	maxOccurs int
}

//initialize the otherAttrs maps and set minOccurs and maxOccurs to 1
func (b *baseParticle) initialize(){
	b.otherAttrs = make(CustomAttributes)
	b.minOccurs = 1
	b.maxOccurs = 1
}

//Base type for Multiple XML Documents and Namespaces elements: import, include and redefine
type baseXmlDocs struct{
	//Required The URI reference to the location of a schema document to include in the target namespace of the containing schema.
	schemaLocation string
	
	//store parse schema from schemaLocation
	schema Schema
}

//container for namespace used in the xsd file
type namespaces struct{
	prefixToUri map[string]string
	uriToPrefix map[string]string
}

//create new namespace object
func New_Namespaces() namespaces{
	return namespaces{
		prefixToUri: make(map[string]string),
		uriToPrefix: make(map[string]string),
	}
}

//Add namespace, return error if prefix or uri is duplicated
func (ns namespaces) add (prefix string, uri string) error{
	if len(ns.prefixToUri[prefix])>0 || len(ns.uriToPrefix[uri])>0 {
		return fmt.Errorf("duplicate namespace found")
	}

	ns.prefixToUri[prefix]=uri
	ns.uriToPrefix[uri]=prefix 
	
	return nil
}

//get uri from the namespaces collection
func (ns namespaces) getUri(prefix string) string{
	return ns.prefixToUri[prefix]
}

//get prefix from the namespaces collection
func (ns namespaces) getPrefix(uri string)string{
	return ns.uriToPrefix[uri]	
} 