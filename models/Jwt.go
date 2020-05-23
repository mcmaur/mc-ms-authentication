package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// DeleteToken : creates a new empty token
func DeleteToken(res http.ResponseWriter, req *http.Request) error {
	http.SetCookie(res, &http.Cookie{
		Name:    "auth_token",
		Value:   "",
		Expires: time.Now().Add(5 * time.Second),
		Domain:  "",
		Path:    "/",
	})
	return nil
}

// CreateToken : creates a new jwt token
func CreateToken(res http.ResponseWriter, req *http.Request, userID uint) error {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return err
	}
	http.SetCookie(res, &http.Cookie{
		Name:    "auth_token",
		Value:   jwtToken,
		Expires: time.Now().Add(120 * time.Minute),
		Domain:  "",
		Path:    "/",
	})
	return nil
}

// TokenValid : check token validity
func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	if tokenString == "" {
		return errors.New("Cookie not found")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

// ExtractToken : extract token info from request
func ExtractToken(r *http.Request) string {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// ExtractTokenID : extract user id info from jwt token
func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
