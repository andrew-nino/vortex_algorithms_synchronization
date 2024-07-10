package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/andrew-nino/vtx_algorithms_synchronization/config"
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"

	repo "github.com/andrew-nino/vtx_algorithms_synchronization/internal/repository/postgresdb"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type tokenClaims struct {
	jwt.StandardClaims
	ManagerId int `json:"user_id"`
}

type AuthService struct {
	repo   repo.Authorization
	config config.Config
}

func NewAuthService(repo repo.Authorization, cfg *config.Config) *AuthService {
	return &AuthService{
		repo:   repo,
		config: *cfg}
}

// Hashes the password and transfers the data to the repository.
func (s *AuthService) CreateManager(mng entity.Manager) (int, error) {
	var err error

	mng.Password, err = generatePasswordHash(mng.Password, s.config)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.CreateManager(mng)
	if err != nil {
		log.Errorf("AuthService.CreateManager - s.repo.CreateManager: %v", err)
		return 0, err
	}

	return id, nil
}

// Checks that the client is already registered and returns the generated token.
func (s *AuthService) SignIn(managerName, password string) (string, error) {

	passwordHash, err := generatePasswordHash(password, s.config)
	if err != nil {
		log.Errorf("AuthService.SignIn - generatePasswordHash: %v", err)
		return "", err
	}
	managerId, err := s.repo.GetManager(managerName, passwordHash)
	if err != nil {
		log.Errorf("AuthService.SignIn - s.repo.GetManager: %v", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(s.config.JWT.TokenTTL)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		managerId,
	})

	return token.SignedString([]byte(s.config.JWT.SigningKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.config.JWT.SigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.ManagerId, nil
}

// generatePasswordHash generates a SHA1 hash of the given password with a salt.
// The salt is a constant string.
func generatePasswordHash(password string, cfg config.Config) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		log.Debugf("failed to generate password hash: %s", err)
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(cfg.JWT.Salt))), nil
}
