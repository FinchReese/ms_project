package jwts

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

func CreateToken(val int, exp time.Duration, secret string, refreshExp time.Duration, refreshSecret string) *JwtToken {
	aExp := time.Now().Add(exp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
	})
	aToken, _ := accessToken.SignedString([]byte(secret))
	rExp := time.Now().Add(refreshExp).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
	})
	rToken, _ := refreshToken.SignedString([]byte(refreshSecret))
	return &JwtToken{
		AccessExp:    aExp,
		AccessToken:  aToken,
		RefreshExp:   rExp,
		RefreshToken: rToken,
	}
}

func ParseToken(tokenString string, secret string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("%v \n", claims)
	} else {
		fmt.Println(err)
	}
}
