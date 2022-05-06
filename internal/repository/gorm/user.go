package gorm

import (
	"context"
)

func (db *Gorm) GetUserKeys(ctx context.Context, userId int) (string, string, error) {
	var user User
	if err := db.Db.WithContext(ctx).Select("PublicKey", "PrivateKey").Where("id = ?", userId).Find(&user).Error; err != nil {
		return "", "", err
	}

	return user.PublicKey, user.PrivateKey, nil
}
