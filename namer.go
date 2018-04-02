package namer

import (
	"github.com/bearchit/namer/analyzer"
	"strings"
	"math/rand"
	"time"
	"github.com/bearchit/namer/pkg/shuffle"
)

type namer struct{
	r *rand.Rand
}

func New() *namer {
	return &namer{
		r: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (n *namer) Name(sentence string, limit int) string {
	vowels, consonants := analyze(sentence)
	vowels = shuffle.String(vowels)
	consonants = shuffle.String(consonants)

	var result string
	for _, c := range consonants {
		result += string(c)
		result += string(vowels[n.r.Int()%len(vowels)])

		if len(result) >= limit {
			break
		}
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