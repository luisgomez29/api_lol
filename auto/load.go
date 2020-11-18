package auto

import (
	"github.com/luisgomez29/api_lol/models"
	"github.com/luisgomez29/api_lol/utils"
	"gorm.io/gorm"
	"log"
)

func Load(db *gorm.DB) {
	mdl := []interface{}{&models.User{}, &models.Character{}}
	err := db.Migrator().DropTable(mdl...)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(mdl...)
	if err != nil {
		log.Fatal(err)
	}

	// Insertar datos de prueba
	db.Create(&users)

	// Mostrar en consola datos insertados
	for _, user := range users {
		utils.Pretty(user)
	}

	db.Create(&characters)
	// Mostrar en consola datos insertados
	for _, product := range characters {
		utils.Pretty(product)
	}
}
