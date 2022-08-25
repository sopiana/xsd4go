package model

//Specifies information to be read or used by users within an annotation element.
// <documentation>
//   source = anyURI
//   xml:lang = language
// Content: ({any})*
// </documentation>
type Documentation struct{
	source string
	lang string
	content string
}