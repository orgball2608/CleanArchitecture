package restaurantstorage

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		// case: error from DB
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
