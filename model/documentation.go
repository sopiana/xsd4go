package model

import (
	"encoding/xml"
	"fmt"
	"strings"
)

//Specifies information to be read or used by users within an annotation element.
// <documentation>
//   source = anyURI
//   xml:lang = language
// Content: ({any})*
// </documentation>
type Documentation struct {
	root *Schema

	//Optional The source of the application information. The source must be a URI reference
	source string

	//Optional The indicator of the language used in the contents.
	lang    string
	content string
}

func New_Documentation(root *Schema) Documentation {
	return Documentation{root: root}
}

func (x Documentation) X(indent uint8) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(`%s<documentation`, getIndent(indent)))

	if len(x.source) > 0 {
		sb.WriteString(fmt.Sprintf(` source="%s"`, x.source))
	}

	if len(x.lang) > 0 {
		sb.WriteString(fmt.Sprintf(` xml:lang="%s"`, x.lang))
	}

	if len(x.content) > 0 {
		sb.WriteString(fmt.Sprintf(`>%s</documentation>`, x.content))
	} else {
		sb.WriteString(`/>`)
	}
	return sb.String()
}

func (x *Documentation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	root := x.root
	for _, attr := range start.Attr {
		switch {
		case attr.Name.Local == "source" && len(attr.Name.Space) == 0:
			x.source = attr.Value
		case attr.Name.Local == "lang" && attr.Name.Space == "http://www.w3.org/XML/1998/namespace":
			x.lang = attr.Value
		default:
			root.Error(d.InputOffset(), "documentation can only have source and xml:lang attributes")
		}
	}
TokenLoop:
	for {
		tok, err := d.Token()
		if err != nil {
			root.Error(d.InputOffset(), "wrong xml syntax")
			return err
		}
		switch token := tok.(type) {
		case xml.StartElement: //ignore inner element
		case xml.CharData:
			x.content += string(token.Copy())
		case xml.EndElement:
			if token.Name == start.Name {
				break TokenLoop
			}
		default:
			return fmt.Errorf("unknown token type on documentation, token: %+v", token)
		}
	}
	return nil
}
