package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var latters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = latters[rand.Intn(len(latters))]
	}

	return string(result)
}
