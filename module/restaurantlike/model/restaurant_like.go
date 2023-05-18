package restaurantlikemodel

import (
	"LearnGo/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

type UserLike struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int `json:"user_id" gorm:"column:user_id;"`
}

func (UserLike) TableName() string {
	return Like{}.TableName()
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func (l *UserLike) GetRestaurantId() int {
	return l.RestaurantId
}
func (l *UserLike) GetUserId() int {
	return l.UserId
}
