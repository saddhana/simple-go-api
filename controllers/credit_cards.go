package controllers

import (
	"net/http"
	"simple-api/models"

	"github.com/gin-gonic/gin"
)

type CreateCreditCardInput struct {
	Number string `json:"number" binding:"required"`
	UserID *uint  `json:"userId"`
}

func FindCreditCards(c *gin.Context) {
	var creditCards []models.CreditCard
	models.DB.Preload("User").Find(&creditCards)

	c.JSON(http.StatusOK, gin.H{"data": creditCards})
}

func CreateCreditCard(c *gin.Context) {
	// Validate input
	var input CreateCreditCardInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creditCard := models.CreditCard{Number: input.Number, UserID: input.UserID}
	models.DB.Create(&creditCard)

	c.JSON(http.StatusOK, gin.H{"data": creditCard})
}
