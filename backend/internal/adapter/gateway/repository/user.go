package repository

import (
	"context"

	"github.com/schema-creator/schema-creator/schema-creator/internal/entities/model"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/herror"
	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/dai"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(gorm *gorm.DB) *UserRepo {
	return &UserRepo{
		db: gorm,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepo) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var (
		user  model.User
		count int64
	)

	result := r.db.Model(&user).Where("user_id = ?", userID).Find(&user).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	if count == 0 {
		return nil, herror.ErrResourceNotFound
	}

	if user.IsDeleted {
		return nil, herror.ErrResourceDeleted
	}

	return &user, nil
}

func (r *UserRepo) GetUsers(ctx context.Context, limit, offset int) ([]model.User, error) {
	var users []model.User
	result := r.db.Model(model.User{}).Where("is_deleted = ?", false).Limit(limit).Offset(offset).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// func (r *UserRepo) UpdateUser(ctx context.Context, userID string, user *model.User) (*model.User, error) {
// 	saveUser := make(map[string]any)
// 	var changeCount int64

// 	if len(user.FirstName) != 0 {
// 		saveUser["first_name"] = user.FirstName
// 		changeCount++
// 	}
// 	if len(user.LastName) != 0 {
// 		saveUser["last_name"] = user.LastName
// 		changeCount++
// 	}
// 	if len(user.FirstNameReading) != 0 {
// 		saveUser["first_name_reading"] = user.FirstNameReading
// 		changeCount++
// 	}
// 	if len(user.LastNameReading) != 0 {
// 		saveUser["last_name_reading"] = user.LastNameReading
// 		changeCount++
// 	}
// 	if user.Age != 0 {
// 		saveUser["age"] = user.Age
// 		changeCount++
// 	}
// 	if user.Gender != 0 {
// 		saveUser["gender"] = user.Gender
// 		changeCount++
// 	}

// 	if changeCount == 0 {
// 		return user, herror.ErrNoChange
// 	}

// 	saveUser["updated_at"] = time.Now()

// 	r.db.Logger = r.db.Logger.LogMode(logger.Info)

// 	result := r.db.Model(&model.User{}).Where("user_id = ?", userID).Updates(saveUser)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	if result.RowsAffected == 0 {
// 		return nil, herror.ErrResourceNotFound
// 	}

// 	return r.GetUserByID(ctx, userID)
// }

// func (r *UserRepo) DeleteUser(ctx context.Context, userID string) error {
// 	result := r.db.Model(&model.User{}).Where("user_id = ?", userID).Updates(map[string]any{
// 		"is_delete": true,
// 	})
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	if result.RowsAffected == 0 {
// 		return herror.ErrResourceNotFound
// 	}

// 	return nil
// }

var _ dai.UserRepo = (*UserRepo)(nil)
