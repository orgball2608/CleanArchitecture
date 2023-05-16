package restaurantlikestorage

import (
	"LearnGo/common"
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) CreateLikeRestaurant(context context.Context, data *restaurantlikemodel.UserLike) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
