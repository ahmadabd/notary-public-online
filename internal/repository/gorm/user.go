package gorm

import (
	"context"
	"notary-public-online/internal/entity/model"

	"gorm.io/gorm"
)

func (db *Gorm) CreateUser(ctx context.Context, userInp *model.User) (model.User, error) {
	user := mapFromUserEntity(*userInp)

	if err := db.Db.WithContext(ctx).Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return MapToUserEntity(user), nil
}

func (db *Gorm) GetUserWithEmail(ctx context.Context, email string) (model.User, error) {

	var user model.User

	if err := db.Db.WithContext(ctx).Where("email = ?", email).Find(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (db *Gorm) CheckUserExistanceWithEmail(ctx context.Context, email string) (bool, error) {

	if err := db.Db.WithContext(ctx).Select("id").Where("email = ?", email).First(&User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *Gorm) GetUserWithId(ctx context.Context, id int) (model.User, error) {
	
	var user model.User

	if err := db.Db.WithContext(ctx).Where("id = ?", id).Find(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
