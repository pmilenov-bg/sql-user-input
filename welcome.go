package screens

import (
	"fmt"
	"pmilenov-bg/sqlite-training/models"
)

func WelcomeScreen(user models.User) {
	fmt.Printf("Welcome, %s!\n", user.Username)
}
