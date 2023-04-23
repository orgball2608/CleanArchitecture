package restaurantbiz

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
)

type DeleteRestaurantStore interface {
	Delete(ctx context.Context, id int) error
	FindDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldDate, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	if oldDate.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.Restaurant{}.TableName(), err)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}
	return nil
}
