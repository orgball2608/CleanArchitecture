package restaurantmodel

import (
	"LearnGo/common"
	"errors"
	"strings"
)

var (
	ErrNameCannotEmpty = errors.New("restaurant name cannot be blank")
)

const EntityName = "Restaurant"

// Business Model
// `common.SQLModel` is embed struct
type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"address" gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	UserId          int                `json:"-" gorm:"column:user_id"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikedCount      int                `json:"liked_count" gorm:"_"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) GetOwnerId() int {
	return r.UserId
}

// Data Model
type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

// Data Model
type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	UserId          int            `json:"-" gorm:"column:user_id"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return ErrNameCannotEmpty
	}
	return nil
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}
