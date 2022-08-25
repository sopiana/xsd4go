package model

import "strings"

const defIndent = "  "

func getIndent(indent uint8) string{
	return strings.Repeat(defIndent, int(indent))
}