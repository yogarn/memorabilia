package jwt

import (
	"errors"
	"log"
	"memorabilia/entity"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	lib_jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Interface interface {
	CreateToken(userID uuid.UUID) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
	GetLoginUser(ctx *gin.Context) (uuid.UUID, error)
}

type jsonWebToken struct {
	SecretKey   string
	ExpiredTime time.Duration
}

type Claims struct {
	UserID uuid.UUID
	lib_jwt.RegisteredClaims
}

func Init() Interface {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Fatalf("can't convert jwt expired time: %v", err)
		return nil
	}
	return &jsonWebToken{
		SecretKey:   secretKey,
		ExpiredTime: time.Duration(expTime) * time.Hour,
	}
}

func (j *jsonWebToken) CreateToken(userID uuid.UUID) (string, error) {
	claim := &Claims{
		UserID: userID,
		RegisteredClaims: lib_jwt.RegisteredClaims{
			ExpiresAt: lib_jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := lib_jwt.NewWithClaims(lib_jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *jsonWebToken) ValidateToken(tokenString string) (uuid.UUID, error) {
	var claim Claims
	var userID uuid.UUID

	token, err := lib_jwt.ParseWithClaims(tokenString, &claim, func(t *lib_jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return userID, err
	}

	if !token.Valid {
		return userID, errors.New("invalid token")
	}

	userID = claim.UserID

	return userID, nil
}

func (j *jsonWebToken) GetLoginUser(ctx *gin.Context) (uuid.UUID, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return uuid.Nil, errors.New("failed to get login user")
	}
	userId := user.(*entity.User).ID
	return userId, nil
}
