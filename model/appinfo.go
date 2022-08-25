package model

import (
	"encoding/xml"
	"fmt"
	"strings"
)

//Specifies information to be used by applications within an annotation element.
// <appinfo
//   source = anyURI>
// Content: ({any})*
// </appinfo>
type AppInfo struct{
	root *Schema

	//Optional. The source of the application information. The source must be a URI reference. 
	source string
	content string
}

func New_AppInfo(root *Schema) AppInfo{
	return AppInfo{root: root}
}

func (x AppInfo) X(indent uint8) string{
	sb:= strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s<appinfo", getIndent(indent)))
	if len(x.source)>0 {
		sb.WriteString(fmt.Sprintf(` source = "%s"`, x.source))
	}
	if len(x.content)>0{
		sb.WriteString(fmt.Sprintf(`>%s</appinfo>`, x.content))
	}else{
		sb.WriteString("/>")
	}
	return sb.String()
}

func (x *AppInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{
	root:=x.root
	
	for _, attr:= range start.Attr{
		if attr.Name.Local=="source" && len(attr.Name.Space)==0{
			x.source = attr.Value
		}else{
			fmt.Println("Masuk sini 1")
			root.Error(d.InputOffset(), "appinfo can only have source attribute")
		}
	}
TokenLoop:
	for{
		tok, err:= d.Token()
		if err!=nil{
			root.Error(d.InputOffset(), "wrong xml syntax")
			return err
		}
		switch token:=tok.(type){
		case xml.StartElement:	//ignore inner element
		case xml.CharData:
			x.content+=string(token.Copy())
		case xml.EndElement:
			if token.Name==start.Name{
				break TokenLoop
			}
		default:
			return fmt.Errorf("unknown token type on appinfo")
		}
	}
	return nil
}