package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Detect takes a string and determines whether it's text or Morse code.
// Converts text to Morse code and vice versa. The output is a string.
func Detect(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("Empty line")
	}

	check := strings.Trim(input, ".- ")

	if check == "" {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}
