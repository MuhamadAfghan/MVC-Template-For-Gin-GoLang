package controllers

import (
	"github.com/gin-gonic/gin"
	"go-health-tracker/models"
	"net/http"
)

func ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func SubmitForm(c *gin.Context) {
	var user models.User

	user.Email = c.PostForm("email")
	user.FullName = c.PostForm("full_name")

	db := models.GetDB()
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
