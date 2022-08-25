package model

//Defines an annotation.
//<annotation
// id = ID 
// {any attributes with non-schema Namespace}...>
// Content: (appinfo | documentation)*
// </annotation> 
type Annotation struct{
	root Schema
	parent Any
	id string
	otherAttrs CustomAttributes
	content[] Any
}