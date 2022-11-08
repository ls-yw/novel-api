package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/woodlsy/woodGin/config"
	"time"
)

//
// CreateJwt
// @Description: 创建jwt
// @param ttl
// @param params
// @return string
//
func CreateJwt(ttl int64, params map[string]interface{}) string {

	claims := jwt.MapClaims{
		"iss": "banzhu9",
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + ttl,
	}

	for key, value := range params {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//var hmacSampleSecret []byte
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(config.Configs.App.JwtSecret))
	return tokenString
}

func ParseJwt(jwtToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Configs.App.JwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}
