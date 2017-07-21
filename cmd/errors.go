package cmd

import (
	"errors"
	"fmt"
)

var ErrSourceLangCodeNotFound = errors.New("The source language code can not be found.")
var ErrTargetLangCodeNotFound = errors.New("The target language code can not be found.")
var ErrNoSourceWord = errors.New("No source word to translate.")

func DealError(err error) {
	fmt.Println(err.Error())
}
