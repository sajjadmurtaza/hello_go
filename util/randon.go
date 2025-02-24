package util

import (
	"fmt"
	"math/rand"
	"time"

	_ "golang.org/x/exp/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString() string {
	result := make([]byte, 6)

	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	fmt.Println(result)

	return string(result)
}

func RandomMoney() int64 {
	return RandomInt(1, 100)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "JPY"}

	fmt.Println(len(currencies))
	return currencies[rand.Intn(len(currencies))]
}
