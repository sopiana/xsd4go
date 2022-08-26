package model

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

// Defines an annotation.
// <annotation
// id = ID
// {any attributes with non-schema Namespace}...>
// Content: (appinfo | documentation)*
// </annotation>
type Annotation struct {
	root   *Schema
	parent Any
	//Optional. The ID of this element. The id value must be of type ID and be unique within the document containing this element.
	id string

	otherAttrs CustomAttributes
	content    []Any
}

func New_Annotation(root *Schema, parent Any) Annotation {
	return Annotation{
		root:       root,
		parent:     parent,
		otherAttrs: make(CustomAttributes),
	}
}

func (x Annotation) X(indent uint8) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(`%s<xs:annotation`, getIndent(indent)))
	if len(x.id) > 0 {
		sb.WriteString(fmt.Sprintf(` id="%s"`, x.id))
	}

	sb.WriteString(x.otherAttrs.X())

	if len(x.content) > 0 {
		sb.WriteString(">")
		for _, c := range x.content {
			sb.WriteString(fmt.Sprintf("\n%s", c.X(indent+1)))
		}
		sb.WriteString(fmt.Sprintf("\n%s</xs:annotation>", getIndent(indent)))
	} else {
		sb.WriteString("/>")
	}
	return sb.String()
}

func (x *Annotation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	root := x.root
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			x.id = attr.Value
			if err := root.addId(attr.Value, x); err != nil {
				root.Error(d.InputOffset(), fmt.Sprint(err))
			}
		} else {
			if err := x.otherAttrs.add(attr, root.namespaces); err != nil {
				root.Error(d.InputOffset(), fmt.Sprint(err))
			}
		}
	}
	var stack Stack
TokenLoop:
	for {
		tok, err := d.Token()
		if err != nil {
			root.Error(d.InputOffset(), "wrong xml syntax")
			return err
		}
		switch token := tok.(type) {
		case xml.StartElement:
			stack.Push(&token)
			if token.Name.Space != "http://www.w3.org/2001/XMLSchema" {
				root.Error(d.InputOffset(), "invalid element name space, allowed element namespace is http://www.w3.org/2001/XMLSchema")
				if err := d.Skip(); err != nil {
					root.Error(d.InputOffset(), "wrong xml syntax")
					return err
				}
			} else {
				switch token.Name.Local {
				case "appinfo":
					child := New_AppInfo(root)
					err = child.UnmarshalXML(d, token)
					x.content = append(x.content, child)
				case "documentation":
					child := New_Documentation(root)
					err = child.UnmarshalXML(d, token)
					x.content = append(x.content, child)
				default:
					root.Error(d.InputOffset(), fmt.Sprintf("element %s is not allowed in annotation, allowed element are appinfo and documentation", token.Name.Local))
				}
				if err != nil {
					return err
				}
			}
		case xml.CharData:
			str := strings.TrimSpace(string(token.Copy()))
			if len(str) > 0 {
				root.Error(d.InputOffset(), "characters are not allowed")
			}
		case xml.EndElement:
			if !stack.IsEmpty() && stack.Top().Name == token.Name {
				stack.Pop()
			}

			if token.Name == start.Name {
				break TokenLoop
			}
		case xml.Comment: //ignore comment
		default:
			return fmt.Errorf("unknown token type on annotation, token: %+v %v", token, reflect.TypeOf(token))
		}
	}
	return nil
}
