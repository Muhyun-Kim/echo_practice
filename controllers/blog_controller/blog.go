package blog_controller

import (
	"log"
	"my-echo-app/controllers/user_controller"
	"my-echo-app/database"
	"my-echo-app/models"
	"net/http"
	"strconv"

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

func DeleteBlog(c echo.Context) error {
	id := c.Param("id")

	blogID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid blog ID",
		})
	}

	blog := models.Blog{}
	if err := database.DB.First(&blog, blogID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Blog not found",
		})
	}

	if err := database.DB.Delete(&blog).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete blog",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Blog deleted successfully",
	})

}

func UpdateBlog(c echo.Context) error {
	id := c.Param("id")

	blogID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid blog ID",
		})
	}

	blog := models.Blog{}
	if err := database.DB.First(&blog, blogID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Blog not found",
		})
	}

	user, err := user_controller.GetUserFromSession(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	if user.ID != blog.AuthorID {
		return c.JSON(http.StatusForbidden, map[string]string{
			"error": "You are not allowed to update this blog",
		})
	}

	updateData := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}
	log.Println(updateData)

	blog.Title = updateData.Title
	blog.Content = updateData.Content

	if err := database.DB.Save(&blog).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update blog",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Blog updated successfully",
	})
}
