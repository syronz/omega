package auth

import (
	"errors"
	"omega/engine"
	"omega/internal/models"
	"omega/pkg/user"
	"omega/utils/password"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Service for injecting user repo
type Service struct {
	Repo   Repo
	Engine engine.Engine
}

// ProvideService for auth is used in wire
func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

// Logout user
func (p *Service) Logout(user Auth) error {
	return p.Repo.Logout(user)
}

// Login User
func (p *Service) Login(auth Auth) (userResult user.User, err error) {
	jwtKey := []byte(p.Engine.Environments.Setting.JWTSecretKey)
	userRepo := user.Repo{Engine: p.Engine}
	if userResult, err = userRepo.FindByUsername(auth.Username); err != nil {
		err = errors.New("Username or Password is wrong")
		return
	}

	if password.Verify(auth.Password, userResult.Password,
		p.Engine.Environments.Setting.PasswordSalt) {

		expirationTime := time.Now().Add(
			time.Duration(p.Engine.Environments.Setting.JWTExpiration) * time.Second)
		claims := &models.JWTClaims{
			Username: auth.Username,
			ID:       userResult.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		var extra struct {
			Token string `json:"token"`
		}
		if extra.Token, err = token.SignedString(jwtKey); err != nil {
			return
		}

		userResult.Extra = extra
		userResult.Password = ""

	} else {
		err = errors.New("Username or Password is wrong")
	}

	return
}
