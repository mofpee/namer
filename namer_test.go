package namer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	n := New()
	name := n.Name("Hello, World", 6)
	assert.NotEmpty(t, name)
	assert.Equal(t, 6, len(name))
}

func TestAnalyze(t *testing.T) {
	v, c := analyze("This is a test sentence")
	assert.Equal(t, "iiaeeee", v)
	assert.Equal(t, "thsststsntnc", c)
}