package restaurantrepo

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	"context"
	"log"
)

type listRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type likeRestaurantStore interface {
	GetRestaurantLikes(
		context context.Context,
		ids []int,
	) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     listRestaurantStore
	likeStore likeRestaurantStore
}

func NewListRestaurantBiz(store listRestaurantStore, likeStore likeRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, likeStore: likeStore}
}

func (biz *listRestaurantRepo) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	likes, err := biz.likeStore.GetRestaurantLikes(context, ids)

	if err != nil {
		log.Print(err)
		return result, err
	}

	for key, value := range result {
		result[key].LikedCount = likes[value.Id]
	}

	if err != nil {
		return nil, err
	}
	return result, nil
}
