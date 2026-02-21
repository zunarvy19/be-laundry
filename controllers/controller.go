package controllers

import (
	"laundry-api/config"
	"laundry-api/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is running"})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	//create user
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//find user
	var foundUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": foundUser.ID,
		"role":    foundUser.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// laundrypackages
func GetPackages(c *gin.Context) {
	var packages []models.LaundryPackage
	if err := config.DB.Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get packages"})
		return
	}
	c.JSON(http.StatusOK, packages)
}

func CreatePackage(c *gin.Context) {
	var pkg models.LaundryPackage
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create package"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "package created successfully"})
}

func UpdatePackage(c *gin.Context) {
	var pkg models.LaundryPackage
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&pkg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update package"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "package updated successfully"})
}

func DeletePackage(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.LaundryPackage{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "package deleted successfully"})
}

// contacts
func GetContacts(c *gin.Context) {
	var contacts []models.Contact
	if err := config.DB.Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get contacts"})
		return
	}
	c.JSON(http.StatusOK, contacts)
}

func CreateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contact created successfully"})
}

func UpdateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contact updated successfully"})
}

func DeleteContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contact deleted successfully"})
}

// webcontent
func GetWebContent(c *gin.Context) {
	var webContent []models.WebContent
	if err := config.DB.Find(&webContent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get web content"})
		return
	}
	c.JSON(http.StatusOK, webContent)
}

func CreateWebContent(c *gin.Context) {
	var webContent models.WebContent
	if err := c.ShouldBindJSON(&webContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&webContent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create web content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "web content created successfully"})
}

func UpdateWebContent(c *gin.Context) {
	var webContent models.WebContent
	if err := c.ShouldBindJSON(&webContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&webContent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update web content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "web content updated successfully"})
}

func DeleteWebContent(c *gin.Context) {
	var webContent models.WebContent
	if err := c.ShouldBindJSON(&webContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&webContent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete web content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "web content deleted successfully"})
}
