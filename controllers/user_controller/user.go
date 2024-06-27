package user_controller

import (
	"my-echo-app/database"
	"my-echo-app/models"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateAccount(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to hash password",
		})
	}

	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user",
		})
	}

	userDTO := models.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return c.JSON(http.StatusCreated, userDTO)
}

func Login(c echo.Context) error {
	loginReq := new(LoginRequest)

	if err := c.Bind(loginReq); err != nil {
		c.Logger().Error("Bind error: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	var user models.User
	if err := database.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		c.Logger().Error("DB error: ", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid email or password",
		})
	}

	if !CheckPasswordHash(loginReq.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid email or password",
		})
	}

	session, ok := c.Get("session").(*sessions.Session)
	if !ok || session == nil {
		c.Logger().Error("Session error: session is nil or not a *sessions.Session")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get session",
		})
	}

	session.Values["email"] = user.Email

	err := session.Save(c.Request(), c.Response())
	if err != nil {
		c.Logger().Error("Session save error: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to save session",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
	})
}

func Logout(c echo.Context) error {
	sessions, ok := c.Get("session").(*sessions.Session)
	if !ok || sessions == nil {
		c.Logger().Error("Session error: session is nil or not a *sessions.Session")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get session",
		})
	}

	sessions.Options.MaxAge = -1
	err := sessions.Save(c.Request(), c.Response())
	if err != nil {
		c.Logger().Error("Session save error: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to invalidate session",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logout successful",
	})
}

func GetProfile(c echo.Context) error {
	email := c.Get("email").(string)

	var userProfile models.User
	if err := database.DB.Where("email = ?", email).First(&userProfile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch user profile",
		})
	}

	return c.JSON(http.StatusOK, userProfile)
}

func GetUserFromSession(c echo.Context) (*models.User, error) {
	session, ok := c.Get("session").(*sessions.Session)
	if !ok || session == nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get session")
	}

	email, ok := session.Values["email"].(string)
	if !ok || email == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "User not logged in")
	}

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user from database")
	}

	return &user, nil
}
