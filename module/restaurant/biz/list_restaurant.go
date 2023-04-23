package restaurantbiz

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
)

type listRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		morekeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store listRestaurantStore
}

func NewListRestaurantBiz(store listRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
