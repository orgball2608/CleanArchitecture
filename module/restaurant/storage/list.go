package restaurantstorage

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName())

	if f := filter; f != nil {
		if f.Status > 0 {
			db = db.Where("status = ?", f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))

	}

	if err := db.
		Limit(int(paging.Limit)).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
