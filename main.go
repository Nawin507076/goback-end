package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/nawin14/course/db"
	"gitlab.com/nawin14/course/handler"
	"gitlab.com/nawin14/course/middleware"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// if err := db.Reset(); err != nil {
	// 	log.Fatal(err)
	// }

	if err := db.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()



	r.Use(cors.Default())


	r.GET("/courses", handler.ListCourses(db))
	r.GET("/courses/:id", handler.GetCourse(db))
	r.POST("/courses", handler.CreateCourse(db))
	r.POST("/classes", handler.CreateClass(db))
	r.POST("/enrollments", middleware.RequireUser(db), handler.EnrollClass(db))
	r.POST("/register", handler.Register(db))
	r.POST("/login", handler.Login(db))

	r.Run(":8000")

}
