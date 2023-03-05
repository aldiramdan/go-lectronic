package seeder

import "lectronic/src/databases/orm/models"

var UserSeed = models.Users{
	{
		Username:     "admin",
		Email:        "admin@email.com",
		Password:     "$2a$10$QmEBoTw3fDlxKEXJyF5QueYXuGC1JWafQ4qpW5kdo25LMq4etj1XW",
		Gender:       "Male",
		Role:         "admin",
		MobileNumber: "08123456789",
		Image:        "https://res.cloudinary.com/duwd9m5ol/image/upload/v1676028039/gorental/default_image.jpg",
		IsActive:     true,
	},
	{
		Username:     "user",
		Email:        "user@email.com",
		Password:     "$2a$10$jR2oS.1dK/TFO7e41cFNEeIx1pXLLKj3ONe7RNzWwdxgX..80jGbC",
		Gender:       "Male",
		Role:         "user",
		MobileNumber: "08123456789",
		Image:        "https://res.cloudinary.com/duwd9m5ol/image/upload/v1676028039/gorental/default_image.jpg",
		IsActive:     true,
	},
}
