package main

import (
	"strings"
	"testing"
)

func TestFindKeyboardDevice(t *testing.T) {
	keyboardExpected := "keyboard"

	keyboardResult, err := findKeyboardDevice()

	if strings.Contains(keyboardResult, keyboardExpected) && err != nil {
		t.Errorf("the keyboardResult isn't the expected: %s", keyboardResult)
	}

}
