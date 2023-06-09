package controllers

import (
	"ORM_MVC/config"
	"ORM_MVC/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

//user

type UserResponse struct {
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

func InitEchoUser() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get users",
			path:                 "/users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"users\":",
		},
	}

	e := InitEchoUser()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get user",
			path:                 "/users/:id",
			userID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"user\":[",
		},
	}

	e := InitEchoUser()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestCreateUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success create user",
			path:                 "/users/",
			userJSON:             `{"name": "User1","email": "User1@gmail.com","password":"12462User1"}`,
			expectStatus:         http.StatusCreated,
			expectBodyStartsWith: "{\"users\":",
		},
	}

	e := InitEchoUser()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, CreateUserController(c)) {
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success delete user",
			path:                 "/users/:id",
			userID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"users\":",
		},
	}

	e := InitEchoUser()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

			if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userID               string
		userJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success update user",
			path:                 "/users/:id",
			userID:               "1",
			userJSON:             `{"Name": "User","Email": "user@gmail.com","password":"12345User"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"users\":",
		},
	}

	e := InitEchoUser()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}

//book

type BookResponse struct {
	Message string        `json:"message"`
	Data    []models.Book `json:"data"`
}

type DeleteBookResponse struct {
	Message string `json:"message"`
}

func InitEchoBook() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get books",
			path:                 "/books",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":[",
		},
	}

	e := InitEchoBook()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetBookController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success get book",
			path:                 "/books/:id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":[",
		},
	}

	e := InitEchoBook()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestCreateBookController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		bookJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success create book",
			path:                 "/books/",
			bookJSON:             `{"judul": "Mentari","penulis": "Lintang","penerbit":"PT.Karya Sastra"}`,
			expectStatus:         http.StatusCreated,
			expectBodyStartsWith: "{\"books\":",
		},
	}

	e := InitEchoBook()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.bookJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, CreateBookController(c)) {
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}

func TestDeleteBookController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		bookID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success delete book",
			path:                 "/books/:id",
			bookID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":",
		},
	}

	e := InitEchoBook()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.bookID)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestUpdateBookController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		bookID               string
		bookJSON             string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success update book",
			path:                 "/books/:id",
			bookID:               "1",
			bookJSON:             `{"judul": "Mentari","penulis": "Lintang","penerbit":"PT.Karya Sastra"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"books\":",
		},
	}

	e := InitEchoBook()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.bookJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.bookID)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartsWith))
		}
	}
}
