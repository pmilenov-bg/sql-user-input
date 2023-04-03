package screens

import (
	"strconv"
)

func RegisterScreen() (string, string, int) {
	email := Prompt("Enter email: ")
	gender := Prompt("Enter you gender: ")
	age_string := Prompt("Enter you age: ")
	age, _ := strconv.Atoi(age_string)

	return email, gender, age
}
