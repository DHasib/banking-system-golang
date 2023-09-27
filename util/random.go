package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//rand.Seed(time.Now().UnixNano())
}

// generqate random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generqate random String of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()

}

// generqate random Owner name
func RandomOwner() string {

	return RandomString(6)
}

// generqate random amount of money
func RandomMoney() int64 {

	return RandomInt(0, 1000)
}

// generqate random amount of money
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "BDT"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
