package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-health-tracker/models"
)

func TestSubmitForm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/submit", SubmitForm)

	dsn := "root:@tcp(localhost:3306)/golang_digiup_ujikom?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"
	models.InitDB(dsn)

	form := url.Values{}
	form.Add("email", "test@example.com")
	form.Add("full_name", "John Doe")

	req, err := http.NewRequest(http.MethodPost, "/submit", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Data saved successfully")
}
