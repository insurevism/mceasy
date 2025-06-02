package middleware

import (
	"mceasy/ent"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	secretKey     string
	secretKeyOnce sync.Once
)

// LoadJWTSecretKey ensures the secret key is loaded only once
func LoadJWTSecretKey() string {
	secretKeyOnce.Do(func() {
		secretKey = viper.GetString("jwt.secret.key")
	})
	return secretKey
}

// Generate JWT Token
func JWTTokenGenerator(user *ent.User) (string, error) {
	secretKey := LoadJWTSecretKey()

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
		"fullname": user.Fullname,
		"avatar":   user.Avatar,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// Validate JWT Token
func JWTTokenValidator(secretKey, tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure token uses the correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
}
