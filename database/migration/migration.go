package migration

import (
	"conduit_api/database"
	"conduit_api/model/entity"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Article{})
	if err != nil {
		fmt.Println("Migration Failed")
		panic(err)
	}
	fmt.Println("Migration Success")
}
