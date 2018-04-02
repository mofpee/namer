package namer

import (
	"github.com/bearchit/namer/analyzer"
	"strings"
	"math/rand"
	"time"
	"github.com/bearchit/namer/pkg/shuffle"
)

func Name(sentence string) string {
	vowels, consonants := analyze(sentence)
	vowels = shuffle.String(vowels)
	consonants = shuffle.String(consonants)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result string
	for _, c := range consonants {
		result += string(c)
		result += string(vowels[r.Int()%len(vowels)])
	}

	return result
}

func analyze(sentence string) (vowels string, consonants string) {
	sentence = strings.ToLower(sentence)
	for _, r := range sentence {
		if analyzer.IsVowel(r) {
			vowels += string(r)
		} else if analyzer.IsConsonant(r) {
			consonants += string(r)
		}
	}
	return
}