package tokenutil

import (
	"math/rand"
	"strconv"
	"strings"
)

func GenerateOTPCode() string {
	sb := strings.Builder{}
	for i := 1; i <= 5; i++ {
		sb.WriteString(strconv.Itoa(rand.Intn(10)))
	}

	return sb.String()
}
