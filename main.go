package main 

import (
    "github.com/gin-gonic/gin"
	"github.com/0kkun/gin-todo-app/infrastructure/database"
	"github.com/0kkun/gin-todo-app/domain/models"
	"os"
	"net/http"
	"strconv"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db := database.DbInit()

    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(c *gin.Context) {
		todos := database.FetchTodo(db)
        c.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
    })

	router.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		rawStatus := c.PostForm("status")
		id, err := strconv.Atoi(rawStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		status := models.Status(id)
		rawTime := c.PostForm("deadline")
		deadline, err := strconv.Atoi(rawTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		todo := models.Todo{Text: text, Status: status, Deadline: deadline}
	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		database.InsertTodo(db, todo)
		c.Redirect(302, "/")
	})

    router.Run(":" + os.Getenv("PORT"))
}