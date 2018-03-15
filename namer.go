package namer

import (
	"github.com/bearchit/namer/analyzer"
	"strings"
	"math/rand"
	"time"
)

func Name(sentence string) string {
	vowels,consontants := analyze(sentence)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result string
	for _, c := range consontants {
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