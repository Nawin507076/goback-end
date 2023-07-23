package handler

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// func sendError(c *gin.Context, status int, err error) {
// 	log.Println(err)
// 	c.JSON(status, gin.H{
// 		"message": err.Error(),
// 	})
// }

const secretKey = "SupperSecret"

type claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func generateToken(userID uint) (string, error) {
	payload := claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "course-api",
		},
	}
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return claim.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (*claims, error) {
	// Key func will verify sig;;n;ing method
	//and return secret Key ;;if ;s;ignning meth;od is math
	//otherwise return errror
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("signing method error")
		}
		return []byte(secretKey), nil
	}
	// Parse Claims
	jwtToken, err := jwt.ParseWithClaims(token, &claims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	//Checking  claims Type
	claims, ok := jwtToken.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// func GetUserByID()
