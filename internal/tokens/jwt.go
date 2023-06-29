package tokens

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/config"
	response "github.com/tekpriest/poprev/internal/types"
)

type TokenData struct {
	UserID string
}

type ReqHeader struct {
	XAuthToken string `reqHeader:"x-auth-token"`
}

type JwtTokenMiddlware interface {
	JwtMiddleware(c *fiber.Ctx) error
}

type jwtTokenMiddlware struct {
	c *config.Config
}

type JwtDataClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func (j *jwtTokenMiddlware) JwtMiddleware(c *fiber.Ctx) error {
	headerToken := c.GetReqHeaders()["X-Auth-Token"]
	if headerToken == "" {
		return response.UnauthorizedResponse(
			c,
			"unauthorized request",
			errors.New("empty authorization token"),
		)
	}

	tokenData, err := j.verifyToken(headerToken)
	if err != nil {
		return response.UnauthorizedResponse(c, err.Error(), nil)
	}

	c.Request().Header.Set("user_id", tokenData.UserID)

	return c.Next()
}

func (j *jwtTokenMiddlware) verifyToken(tk string) (*TokenData, error) {
	var tokenData TokenData
	secret := []byte(j.c.JwtSecret)

	token, err := jwt.Parse(tk, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("invalid sining method: %v", t.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if err := token.Claims.Valid(); err != nil {
		return nil, fmt.Errorf("token is expired: %v", err)
	}
	claims := token.Claims.(jwt.MapClaims)
	tokenData.UserID, _ = claims["user_id"].(string)

	return &tokenData, nil
}

func NewJwtMiddleware(c *config.Config) JwtTokenMiddlware {
	return &jwtTokenMiddlware{c: c}
}
