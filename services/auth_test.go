package services

import (
    "testing"
    "time"
    "github.com/golang-jwt/jwt/v4"
)

func TestGenerateJWT(t *testing.T) {
    username := "testuser"

    token, err := GenerateJWT(username)
    if err != nil {
        t.Fatalf("expected no error, but got %v", err)
    }

    if token == "" {
        t.Fatalf("expected token, but got empty string")
    }
}

func TestValidateJWT(t *testing.T) {
    username := "testuser"

    token, err := GenerateJWT(username)
    if err != nil {
        t.Fatalf("expected no error, but got %v", err)
    }

    validatedUsername, err := ValidateJWT(token)
    if err != nil {
        t.Fatalf("expected no error, but got %v", err)
    }

    if validatedUsername != username {
        t.Fatalf("expected %s, but got %s", username, validatedUsername)
    }
}

func TestExpiredToken(t *testing.T) {
    jwtKey = []byte("my_secret_key")

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": "testuser",
        "exp":      time.Now().Add(-time.Hour).Unix(),
    })

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        t.Fatalf("failed to sign token: %v", err)
    }

    _, err = ValidateJWT(tokenString)
    if err == nil {
        t.Fatalf("expected an error, but got nil")
    }

    if ve, ok := err.(*jwt.ValidationError); ok {
        if ve.Errors != jwt.ValidationErrorExpired {
            t.Fatalf("expected validation error for expired token, but got %v", ve)
        }
    } else {
        t.Fatalf("expected jwt validation error, but got %v", err)
    }
}
