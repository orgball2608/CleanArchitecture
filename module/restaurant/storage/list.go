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
	morekeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}

		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(paging.FakeCursor)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
