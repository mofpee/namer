package shuffle

import "math/rand"

func String(v string) string {
	sv := []byte(v)
	rand.Shuffle(len(sv), func(i, j int) {
		sv[i], sv[j] = sv[j], sv[i]
	})
	return string(sv)
}