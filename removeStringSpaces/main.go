package main

import "fmt"

func main() {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test with spaces", " dsoijf osidjf osdij fois  so dijfo  osdij ", "dsoijfosidjfosdijfoissodijfoosdij"},
		{"Empty string", "", ""},
		{"No spaces", "nospace", "nospace"},
		{"Multiple spaces", "this  is   a test", "thisisatest"},
		{"Special characters", "!@#$%^&*()  ", "!@#$%^&*()"},
	}

	allTestsPassed := true

	for _, test := range tests {
		result := noSpace(test.input)
		if result != test.expected {
			fmt.Println(test.name, "failed:", result)
			allTestsPassed = false
		}
	}

	if allTestsPassed {
		fmt.Println("All tests passed")
	}
}

func noSpace(givenString string) string {
	var modifiedString string

	for _, char := range givenString {
		if char != ' ' {
			modifiedString += string(char)
		}
	}

	return modifiedString
}
