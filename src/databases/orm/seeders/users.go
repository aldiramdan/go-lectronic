package seeder

import "lectronic/src/databases/orm/models"

var UserSeed = models.Users{
	{
		Username:     "admin",
		Email:        "admin@email.com",
		Password:     "$2a$12$LzPmxfEZoVbCpGUGticqreZHbKLJICuXHPjwOPMZ9OFrmSDWHyPQW",
		Gender:       "Male",
		Role:         "admin",
		MobileNumber: "08123456789",
		Image:        "https://res.cloudinary.com/duwd9m5ol/image/upload/v1676028039/gorental/default_image.jpg",
	},
	{
		Username:     "user",
		Email:        "user@email.com",
		Password:     "$2a$12$wcGtHuywxUX8fvxYqv8aJ.A0JcasSMqFglWtoIjNQxNRQlPQ/ChGO",
		Gender:       "Male",
		Role:         "user",
		MobileNumber: "08123456789",
		Image:        "https://res.cloudinary.com/duwd9m5ol/image/upload/v1676028039/gorental/default_image.jpg",
	},
}
