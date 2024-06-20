package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

// getJwtToken generates a JWT for a given user using their password as the key
func getJwtToken(user *User) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the user details and expiry time
	claims := &Claims{
    ID: user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Email:     user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "my-auth-server",
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	// Declare the token with the algorithm used for signing and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Use the user's password as the key for signing the token
	tokenString, err := token.SignedString([]byte(user.Password))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// decodeJwtToken decodes and validates a JWT token using the user's password as the key
func decodeJwtToken(tokenString string, user User) (*Claims, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// `token` is the parsed token, `err` is the error
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(user.Password), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token signature")
		}
		return nil, fmt.Errorf("invalid token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

