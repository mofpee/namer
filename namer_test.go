package namer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamer_NameDefault(t *testing.T) {
	n := New("default")
	name := n.Name("Hello, World", 6)
	assert.NotEmpty(t, name)
}

func TestNamer_NameMixedShuffle(t *testing.T) {
	n := New("mixedshuffle")
	name := n.Name("Hello, World", 6)
	assert.NotEmpty(t, name)
	assert.Equal(t, 6, len(name))
}

func TestNamer_NameInvalid(t *testing.T) {
	n := New("unknown")
	name := n.Name("Hello, World", 6)
	assert.NotEmpty(t, name)
}

func TestAnalyze(t *testing.T) {
	v, c := analyze("This is a test sentence")
	assert.Equal(t, "iiaeeee", v)
	assert.Equal(t, "thsststsntnc", c)
}