package auto

import "github.com/luisgomez29/api_lol/models"

var (
	users = []models.User{
		{FirstName: "Luis", LastName: "Gómez", Email: "luis.gomez@usantoto.edu.co", Password: "123456"},
	}

	characters = []models.Character{
		{Name: "Akali", Position: "Central"},
		{Name: "Lucian", Position: "Tirador "},
		{Name: "Graves", Position: "Jungla"},
		{Name: "Camille", Position: "Superior "},
	}
)
