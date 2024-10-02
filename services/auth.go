package services

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

func GenerateJWT(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 2).Unix(),
    })
    
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ValidateJWT(tokenString string) (interface{}, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
        }
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["username"], nil
    } else {
        return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorExpired)
    }
}
