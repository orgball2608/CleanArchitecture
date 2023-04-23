package restaurantbiz

import (
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
	UpdateData(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
	FindDataWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(context context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate) error {

	result, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(context, id, data); err != nil {
		return err
	}
	return nil
}
