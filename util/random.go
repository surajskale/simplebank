package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// initialization
func init() {
	rand.Seed(time.Now().UnixNano())
	// r := rand.New(rand.NewSource(SEED))
	// fmt.Println(r.Uint64())
	// fmt.Println(r.Uint64())

}

// RandomInt generates random integer in between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(n)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates random owner name of fixed size
func RandomOwner() string {
	OwnerNameLength := 6
	return RandomString(OwnerNameLength)
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "INR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
