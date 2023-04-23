package restaurantbiz

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
)

type GetRestaurantStore interface {
	FindDataWithCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(
	ctx context.Context,
	id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.Restaurant{}.TableName(), err)

		}
		// OR: able to throw err `sth went wrong with server`
		return nil, common.ErrCannotGetEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	// for case soft deleted (mean: can't retrieve record when status == 0)
	if result.Status == 0 {
		// FOR CASE Security:
		return nil, common.ErrEntityDeleted(restaurantmodel.Restaurant{}.TableName(), err)
	}
	return result, err
}
