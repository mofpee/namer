package analyzer

var (
	alphabets string
	vowels = "aeiouwy"
	consonants string
)

func init() {
	for i:=97; i<=122; i++ {
		alphabets += string(i)
	}

	for _, r := range alphabets {
		if !contains(r, vowels) {
			consonants += string(r)
		}
	}
}

func contains(v rune, runes string) bool {
	for _, r := range runes {
		if v == r {
			return true
		}
	}
	return false
}

func IsVowel(r rune) bool {
	return contains(r, vowels)
}

func IsConsonant(r rune) bool {
	return contains(r, consonants)
}