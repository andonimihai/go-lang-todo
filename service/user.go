package service

import (
	"go-gin-todo/entity"
)

func UpsertUser(userInfo entity.UpsertUser) (entity.User, error) {
	var user entity.User
	result := entity.DB.Where(entity.User{ExternalId: userInfo.UserId}).Attrs(entity.User{
		Name:       userInfo.Name,
		Email:      userInfo.Email,
		ExternalId: userInfo.UserId,
		ImageUrl:   userInfo.Avatar,
	}).FirstOrCreate(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil

}
