package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//TokenPayload date saved in token
type TokenPayload struct {
	UserUUID string
}

//TokenDetails details of token
type TokenDetails struct {
	AccessToken string
	AtExpires   int64
}

//Token struct for assign interface
type Token struct{}

//NewToken function to create a new instance
func NewToken() *Token {
	return &Token{}
}

//TokenInterface for sign token struct
type TokenInterface interface {
	CreateToken(payload TokenPayload) (*TokenDetails, error)
	CheckToken(tokenString string) (*TokenPayload, error)
}

//Token implements the TokenInterface
var _ TokenInterface = &Token{}

//CreateToken create a new token
func (t *Token) CreateToken(payload TokenPayload) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Second * 10).Unix()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = td.AtExpires
	atClaims["user_uuid"] = payload.UserUUID

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

//TokenValid check if token is valid
func TokenValid(tokenString string) error {
	verifiedToken, err := VerifyToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := verifiedToken.Claims.(jwt.Claims); !ok && !verifiedToken.Valid {
		return err
	}
	return nil
}

//VerifyToken is ok
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//CheckToken responsible for check token
func (t *Token) CheckToken(tokenString string) (*TokenPayload, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		UserUUID, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, err
		}
		return &TokenPayload{
			UserUUID: UserUUID,
		}, nil
	}
	return nil, err
}
