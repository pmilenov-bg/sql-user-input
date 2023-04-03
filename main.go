package main

import (
	"database/sql"
	"fmt"
	"log"
	"pmilenov-bg/sqlite-training/database"
	"pmilenov-bg/sqlite-training/models"
	"pmilenov-bg/sqlite-training/screens"
	"pmilenov-bg/sqlite-training/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sql.DB
	db = database.OpenDatabase()
	defer db.Close()
	database.CreateTables(db)

	var user models.User
	user.Username = screens.Prompt("Enter username: ")
	user.Email = screens.Prompt("Enter email: ")

	// Check if the user already exists in the database
	err := db.QueryRow("SELECT id, username, email FROM users WHERE username = ? and email = ?", user.Username, user.Email).Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User does not exist, creating a new account.")
			user.Email, user.Gender, user.Age = screens.RegisterScreen() // user.Email, user.Gender, user.Age = RegisterScreen()
			services.DoRegister(db, user)

		} else {
			log.Fatal(err)
		}
	}

	screens.WelcomeScreen(user)

	for {
		choice := screens.Prompt(`
		1 - add product
		2 - list products
		-----
		`)
		switch choice {
		case "1":
			for {
				product := models.Product{User_id: user.Id}
				screens.DefineProduct(&product)

				// PETE: product, err = product.Insert(db)
				product, err = models.InsertProduct(db, product)
				if err != nil {
					log.Fatal(err)
				}
				// Check if user wants to add another product
				addAnother := screens.Prompt("Add another product? (y/n)")
				if addAnother != "y" {
					break
				}
			}
		case "2":
			fmt.Println("wip...")
			// product := models.GetProduct()
			// select by 	1 - brands' names
			// 				2 - type
			// 				3 - user_id
			// press something to exit

		default:
			// handle invalid input
		}
	}
	// for {
	// 	choice := screens.Prompt(`
	// 	1 - add product
	// 	2 - list products
	// 	-----
	// 	`)
	// 	switch choice {
	// 	case "1":
	// 		product := models.GetProduct{
	// 			brand: screens.Prompt("What brand?"),
	// 		}
	// 		product, err = addProduct(db, product, user)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	case "2":

	// 	default:

	// 	}
	// }

}
