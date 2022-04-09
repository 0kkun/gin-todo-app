package database

import (
	"fmt"
	"github.com/0kkun/gin-todo-app/domain/models"
	"github.com/jinzhu/gorm"
)

// Todoのマイグレート
func DbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.AutoMigrate(&models.Todo{})
	return db
}

// DBの作成
func DbCreate(todo models.Todo) {
	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	db.Create(&todo)
}

func DbRead(id ...int) []models.Todo {
	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todos []models.Todo
	db.Find(&todos)
	return todos
}

func DbUpdate(id int, text string, status models.Status, deadline int) models.Todo {
	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
	if err != nil {
		fmt.Errorf("could not open database")
	}
	var todo models.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Deadline = deadline
	db.Save(&todo)
	return todo
}

func DbDelete(id int) {
	db, err := gorm.Open("mysql", "root:@tcp(locakhost:3306)/github.com/0kkun/gin-todo-app?parseTime=true")
	if err != nil {
		fmt.Errorf("could not open database")
	}
    var todo models.Todo
    db.First(&todo, id)
    db.Delete(&todo)
}