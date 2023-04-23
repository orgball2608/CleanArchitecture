package restaurantstorage

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
)

func (s *sqlStore) UpdateData(
	context context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
