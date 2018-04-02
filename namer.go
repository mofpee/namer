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
	fn namerFunc
}

type namerFunc func(string,int) string

func New(algorithm string) *namer {
	n:= &namer{
		r: rand.New(rand.NewSource(time.Now().Unix())),
	}

	switch algorithm {
	case "mixedshuffle":
		n.fn = n.NameMixedShuffle

	case "default":
		fallthrough
	default:
		n.fn = n.NameDefault
	}

	return n
}

func (n *namer) Name(sentence string, limit int) string {
	return n.fn(sentence,limit)
}

func (n *namer) NameDefault(sentence string, limit int) string {
	vowels, consonants := analyze(sentence)
	s := shuffle.New(n.r)
	vowels = s.String(vowels)
	consonants = s.String(consonants)

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

func (n *namer) NameMixedShuffle(sentence string, limit int) string {
	vowels, consonants := analyze(sentence)
	s := shuffle.New(n.r)
	return s.String(vowels + consonants)[:limit]
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

