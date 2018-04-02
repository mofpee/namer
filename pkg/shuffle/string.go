package shuffle

import "math/rand"

type shuffle struct {
	r *rand.Rand
}

func New(r *rand.Rand) *shuffle {
	return &shuffle{
		r: r,
	}
}

func (s *shuffle) String(v string) string {
	sv := []byte(v)
	s.r.Shuffle(len(sv), func(i, j int) {
		sv[i], sv[j] = sv[j], sv[i]
	})
	return string(sv)
}