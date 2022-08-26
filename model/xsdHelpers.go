package model

import (
	"encoding/xml"
	"strings"
)

const defIndent = "  "

func getIndent(indent uint8) string {
	return strings.Repeat(defIndent, int(indent))
}

type Stack []*xml.StartElement

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(elem *xml.StartElement) {
	*s = append(*s, elem)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (*xml.StartElement, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Top() *xml.StartElement {
	if s.IsEmpty() {
		return nil
	} else {
		index := len(*s) - 1
		return (*s)[index]
	}
}
