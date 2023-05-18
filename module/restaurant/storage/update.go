package restaurantstorage

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
	"gorm.io/gorm"
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

func (s *sqlStore) IncreaseLikeCount(context context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
func (s *sqlStore) DecreaseLikeCount(context context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
