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
	store     DeleteRestaurantStore
	requester common.Requester
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore, requester common.Requester) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store, requester: requester}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.Restaurant{}.TableName(), err)
	}

	if oldData.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.Restaurant{}.TableName(), err)
	}
	return nil
}
