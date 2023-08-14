package utils

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// // Random int between min and max
func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
