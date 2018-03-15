package analyzer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testVowels = "aeiouwy"
	testConsonants = "bcdfghjklmnpqrstvxz"
)

func TestInit(t *testing.T) {
	assert.Equal(t, 26, len(alphabets))
	assert.Equal(t, 7, len(vowels))
	assert.Equal(t, 26-7, len(consonants))
}

func TestIsVowel(t *testing.T) {
	for _, r := range testVowels {
		assert.True(t, IsVowel(r))
	}

	for _, r := range testConsonants {
		assert.False(t, IsVowel(r))
	}

	assert.False(t, IsVowel(' '))
}

func TestIsConsonant(t *testing.T) {
	for _, r := range testVowels {
		assert.False(t, IsConsonant(r))
	}

	for _, r := range testConsonants {
		assert.True(t, IsConsonant(r))
	}

	assert.False(t, IsVowel(' '))
}