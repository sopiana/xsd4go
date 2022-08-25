package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIndent(t *testing.T) {
	assert.Equal(t, "", getIndent(0))
	assert.Equal(t, "  ", getIndent(1))
	assert.Equal(t, "        ", getIndent(4))
}