package main

import (
	"golang-test/delivery"
	"golang-test/models"
	"golang-test/repository"
	"golang-test/usecase"
	"net/http"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conn := "root:@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.Bank{})
	db.AutoMigrate(&models.Merchant{})
	db.AutoMigrate(&models.Customer{})
	newRepositories := repository.NewRepositories(db)
	newUseCase := usecase.NewUseCase(newRepositories)
	newDelivery := delivery.NewDelivery(newUseCase)

	router := gin.Default()
	api := router.Group("/api")
	api.POST("/customers/register", authMiddleWare(newUseCase), newDelivery.RegisterCustomer)
	api.POST("/merchants/register", newDelivery.RegisterMerchant)
	api.POST("/auth/login", newDelivery.Login)
	api.POST("/transaction", newDelivery.Transaction)

	router.Run()
}

func authMiddleWare(authCustomer usecase.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.Contains(header, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized1")
			return
		}
		token := ""
		arrayToken := strings.Split(header, " ")
		token = arrayToken[1]
		validateToken, err := authCustomer.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized2")
			return
		}

		claim, ok := validateToken.Claims.(jwt.MapClaims)
		if !ok || !validateToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized3")
			return
		}

		userID := int(claim["user_id"].(float64))
		customer, err := authCustomer.GetCustomerById(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized4")
			return
		}

		c.Set("currentUser", customer)

	}
}
