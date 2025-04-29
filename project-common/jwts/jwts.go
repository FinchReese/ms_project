package jwts

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

func CreateToken(val int64, exp time.Duration, secret string, refreshExp time.Duration, refreshSecret string, ip string) *JwtToken {
	aExp := time.Now().Add(exp).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
		"ip":    ip,
	})
	aToken, _ := accessToken.SignedString([]byte(secret))
	rExp := time.Now().Add(refreshExp).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": val,
		"exp":   aExp,
		"ip":    ip,
	})
	rToken, _ := refreshToken.SignedString([]byte(refreshSecret))
	return &JwtToken{
		AccessExp:    aExp,
		AccessToken:  aToken,
		RefreshExp:   rExp,
		RefreshToken: rToken,
	}
}

func ParseToken(tokenString string, secret string, ip string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		zap.L().Error("jwt.Parse error", zap.Error(err), zap.String("tokenString", tokenString))
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		val := int64(claims["token"].(float64))
		exp := int64(claims["exp"].(float64))
		if exp < time.Now().Unix() {
			return 0, errors.New("token过期")
		}
		if ip != claims["ip"] {
			return 0, errors.New("ip不匹配")
		}
		return val, nil
	}
	return 0, fmt.Errorf("token invalid")
}
