package ucase

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/xtabs12/test_dans/internal/dto"
	"github.com/xtabs12/test_dans/internal/repo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type authImpl struct {
	userRepo repo.User
}

func NewAuth(userRepo repo.User) AUTH {
	return &authImpl{
		userRepo: userRepo,
	}
}

type AUTH interface {
	Authenticate(ctx context.Context, params dto.AuthenticateParams) (*dto.Token, error)
}

func (i *authImpl) Authenticate(ctx context.Context, params dto.AuthenticateParams) (*dto.Token, error) {

	user, getUserErr := i.userRepo.FindUserByUserName(ctx, params.Username)
	if getUserErr != nil {
		log.Println(getUserErr)
		return nil, getUserErr
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		log.Println(err)
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": map[string]interface{}{
			"id":       user.ID,
			"username": user.UserName,
			"name":     user.Name,
		},
		"timestamps": time.Now().Unix(),
		"exp":        time.Now().Add(30 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &dto.Token{
		TokenString: tokenString,
	}, nil
}
