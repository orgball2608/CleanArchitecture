package restaurantmodel

import "LearnGo/common"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}
