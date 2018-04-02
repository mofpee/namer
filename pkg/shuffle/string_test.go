package shuffle

import (
	"testing"
	"math/rand"
	"time"
)

func TestString(t *testing.T) {
	s := New(
		rand.New(rand.NewSource(time.Now().Unix())),
	)
	t.Log(s.String("Hello"))
}
