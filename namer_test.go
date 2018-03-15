package namer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	name := Name("Hello, World")
	assert.NotEmpty(t, name)
}

func TestAnalyze(t *testing.T) {
	v, c := analyze("This is a test sentence")
	assert.Equal(t, "iiaeeee", v)
	assert.Equal(t, "thsststsntnc", c)
}

