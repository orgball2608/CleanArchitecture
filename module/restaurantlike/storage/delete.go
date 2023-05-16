package restaurantlikestorage

import (
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Delete(context context.Context, userId, restaurantId int) error {
	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
