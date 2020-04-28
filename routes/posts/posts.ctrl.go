package posts

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cjaewon/echo-gorm-example/database/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Read : Read Router
func (PostsRouter) Read(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusPreconditionFailed)
	}

	db, _ := c.Get("db").(*gorm.DB)

	var post models.Post
	if err := db.Where("id = ?", id).First(&post).Error; err == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, post)
}

// Create : Create Router
func (PostsRouter) Create(c echo.Context) error {
	type RequestBody struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db, _ := c.Get("db").(*gorm.DB)
	username, _ := c.Get("username").(string)

	post := models.Post{
		Title:    body.Title,
		Content:  body.Content,
		Date:     time.Now(),
		Username: username,
	}

	db.Create(post)

	return c.JSON(http.StatusOK, post)
}

// Delete : Delete Router
func (PostsRouter) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusPreconditionFailed)
	}

	db, _ := c.Get("db").(*gorm.DB)

	var post models.Post
	if err := db.Where("id = ?", id).First(&post).Error; err == nil {
		return c.NoContent(http.StatusNotFound)
	}

	db.Delete(&post)

	return c.NoContent(http.StatusOK)
}
