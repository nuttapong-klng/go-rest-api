package main

import (
	"go_rest_api/authentication"
	"go_rest_api/user"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=go_rest_api dbname=go_rest_api host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	userAPI := user.Handlers{
		UserRepo: user.RepositoryPg{
			DBConnection: db,
		},
	}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		public := v1.Group("/")
		{
			user := public.Group("user")
			{
				user.POST("registration", userAPI.UserRegistrationHandlers)
				user.POST("login", userAPI.UserLoginHandlers)
			}
		}
		authorized := r.Group("/")
		authorized.Use(authentication.AuthorizeJWT())
		{

		}
	}
	r.Run(":8000")
}
