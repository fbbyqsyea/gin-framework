package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func RandString(l int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, l)
	rand.Read(b) //整合
	return hex.EncodeToString(b)[:l]
}
