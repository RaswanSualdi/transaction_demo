package main

import (
	"golang-test/delivery"
	"golang-test/models"
	"golang-test/repository"
	"golang-test/usecase"
	"net/http"
	"strings"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	conn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
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
	api.POST("/customers/register", newDelivery.RegisterCustomer)
	api.POST("/merchants/register", newDelivery.RegisterMerchant)
	api.POST("/auth/login", newDelivery.Login)
	api.POST("/transaction", authMiddleWare(newUseCase), newDelivery.Transaction)

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
