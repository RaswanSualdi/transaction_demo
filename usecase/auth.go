package usecase

import (
	"errors"
	"golang-test/models"
	"golang-test/payload/request"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("SECRET_KEY")

func (u *usecase) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["iat"] = time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (u *usecase) ValidateToken(Encodedtoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(Encodedtoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid Token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (u *usecase) Login(input request.Login) (models.Customer, error) {
	username := input.Username
	password := input.Password
	customer, err := u.repositories.FindCustomerByEmail(username)
	if err != nil {
		return customer, err
	}

	if customer.ID == 0 {
		return customer, errors.New("customer doesn't found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(password))

	if err != nil {
		return customer, err
	}

	return customer, nil

}

func (u *usecase) GetCustomerLogin(c *gin.Context) (models.Customer, error) {
	header := c.GetHeader("Authorization")
	token := ""
	arrayToken := strings.Split(header, " ")
	token = arrayToken[1]
	validateToken, _ := u.ValidateToken(token)
	claim, _ := validateToken.Claims.(jwt.MapClaims)

	userID := int(claim["user_id"].(float64))

	customer, err := u.GetCustomerById(userID)
	if err != nil {
		return customer, err
	}

	return customer, nil

}
