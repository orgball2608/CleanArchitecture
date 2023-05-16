package restaurantlikestorage

import (
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantlikemodel.Like, error) {
	var data restaurantlikemodel.Like

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
