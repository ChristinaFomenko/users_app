package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64() //nolint:gosec // not changing math/rand
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(min-max+1) //nolint:gosec // not changing math/rand
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] //nolint:gosec // not changing math/rand
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6) //nolint:gomnd // not
}

func RandomMoney() float64 {
	return RandomFloat(0.01, 999999.99) //nolint:gomnd // not
}

func RandomYear() int64 {
	return RandomInt(1900, 2021) //nolint:gomnd // not
}
