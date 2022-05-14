package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	tmp := make([]rune, len(str))
	for index, str := range str {
		tmp[index] = str
	}

	lenTmp := len(tmp)
	for i := 0; i < lenTmp; i++ {
		if unicode.IsDigit(tmp[i]) {
			return "", ErrInvalidString
		}
		if i+1 == lenTmp {
			result.WriteRune(tmp[i])
		} else {
			n, err := strconv.Atoi(string(tmp[i+1]))
			if err == nil {
				result.WriteString(strings.Repeat(string(tmp[i]), n))
				i++
			} else {
				result.WriteRune(tmp[i])
			}
		}
	}
	return result.String(), nil
}
