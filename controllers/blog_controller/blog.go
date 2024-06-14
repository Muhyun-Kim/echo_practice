package blog_controller

import (
	"my-echo-app/database"
	"my-echo-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBlog(c echo.Context) error {
	blog := new(models.Blog)

	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if err := database.DB.Create(&blog).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create blog",
		})
	}

	if err := database.DB.Preload("Author").First(&blog, blog.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	authorDTO := models.AuthorDTO{
		ID:        blog.Author.ID,
		Username:  blog.Author.Username,
		Email:     blog.Author.Email,
		CreatedAt: blog.Author.CreatedAt,
	}

	blogDTO := models.BlogDTO{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		AuthorID:  blog.AuthorID,
		Author:    authorDTO,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
		Comments:  []models.CommentDTO{},
	}

	return c.JSON(http.StatusOK, blogDTO)
}

func GetBlogs(c echo.Context) error {
	blogs := []models.Blog{}
	if err := c.Bind(blogs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if err := database.DB.Preload("Author").Find(&blogs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve blogs",
		})
	}

	var blogDTOs []models.BlogDTO
	for _, blog := range blogs {
		authorDTO := models.AuthorDTO{
			ID:        blog.Author.ID,
			Username:  blog.Author.Username,
			Email:     blog.Author.Email,
			CreatedAt: blog.Author.CreatedAt,
		}
		blogDTO := models.BlogDTO{
			ID:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			AuthorID:  blog.AuthorID,
			Author:    authorDTO,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
			Comments:  []models.CommentDTO{},
		}
		blogDTOs = append(blogDTOs, blogDTO)
	}
	return c.JSON(http.StatusOK, blogDTOs)
}
