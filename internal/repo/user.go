package repo

import (
	"context"
	"github.com/xtabs12/test_dans/internal/entity"

	"gorm.io/gorm"
)

type userImpl struct {
	db *gorm.DB
}
type User interface {
	FindUserByUserName(ctx context.Context, username string) (*entity.User, error)
}

func NewUser(db *gorm.DB) User {
	return &userImpl{db: db}
}

func (i *userImpl) FindUserByUserName(ctx context.Context, username string) (*entity.User, error) {

	var mod entity.User
	if err := i.db.
		Debug().
		Model(mod).
		WithContext(ctx).
		Where("username = ?", username).
		First(&mod).Error; err != nil {
		return nil, err
	}

	return &mod, nil
}
