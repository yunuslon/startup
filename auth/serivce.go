package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("ST4RTup_s3cr3t_k3y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signdToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signdToken, err
	}

	return signdToken, nil
}
