package cipher

import (
	"bytes"
	"math"
	"strings"
)

// Rot13 ...
func Rot13(msg string) string {

	result := make([]string, 0)
	key := 13
	upperSym := []byte{
		65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
		81, 82, 83, 84, 85, 86, 87, 88, 89, 90}

	for _, val := range msg {

		// convert lower to upper case
		if val >= 97 && val < 123 {
			val = val ^ 32
		}

		idx := bytes.IndexRune(upperSym, val)

		if idx == -1 {
			result = append(result, string(val))
		} else {

			idx += key

			// if over go back
			if idx >= len(upperSym) {
				idx = int(math.Abs(float64(idx - len(upperSym))))
			}

			result = append(result, string(upperSym[idx]))
		}
	}

	return strings.Join(result, "")
}
