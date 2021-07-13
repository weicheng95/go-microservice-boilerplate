package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"
	"gitlab.com/kiplexlab/go-microservice-boilerplate/controllers"
	"gitlab.com/kiplexlab/go-microservice-boilerplate/db"
	"gitlab.com/kiplexlab/go-microservice-boilerplate/forms"
	"gitlab.com/kiplexlab/go-microservice-boilerplate/middleware"
	"gitlab.com/kiplexlab/go-microservice-boilerplate/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	//Start the default gin server
	r := gin.Default()

	//Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	//Start MYSQL database
	//Example: db.GetDB() - More info in the models folder
	db.Init()

	// Migrate the schema
	db.GetDB().AutoMigrate(&models.Todo{})

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		todo := new(controllers.TodoController)

		v1.POST("/todo", todo.Create)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ginBoilerplateVersion": "v1.0.0",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
	})

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	r.Run(":" + port)
}
