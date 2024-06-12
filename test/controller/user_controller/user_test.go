package usercontroller_test

import (
	"my-echo-app/database"
	testdatabase "my-echo-app/test/test_database"
	testset "my-echo-app/test/test_set"
	"testing"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"my-echo-app/controllers/user_controller"
	"my-echo-app/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testset.SetupTestMain(m)
}

func TestCreateAccount(t *testing.T) {
	e := echo.New()

	user := models.User{
		Email:    "test@example.com",
		Password: "password",
	}
	payload, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if database.DB == nil {
		t.Fatalf("database.DB is nil")
	}
	if testdatabase.Mock == nil {
		t.Fatalf("testdatabase.Mock is nil")
	}

	testdatabase.Mock.ExpectBegin()
	testdatabase.Mock.ExpectExec("INSERT INTO `users`").
		WithArgs(sqlmock.AnyArg(), user.Email, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	testdatabase.Mock.ExpectCommit()

	if assert.NoError(t, user_controller.CreateAccount(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var responseUser models.User
		if err := json.Unmarshal(rec.Body.Bytes(), &responseUser); err != nil {
			t.Errorf("expected response body to be valid JSON, got '%s'", rec.Body.String())
		} else {
			assert.Equal(t, user.Email, responseUser.Email)
		}
	}
}
