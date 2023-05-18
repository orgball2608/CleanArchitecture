package restaurantlikebussiness

import (
	"LearnGo/common"
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"context"
	"log"
)

type UserDisLikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(context context.Context, id int) error
}

type userDisLikeRestaurantBiz struct {
	store    UserDisLikeRestaurantStore
	decStore DecreaseLikeCountStore
}

func NewUserDisLikeRestaurantBiz(
	store UserDisLikeRestaurantStore, decStore DecreaseLikeCountStore) *userDisLikeRestaurantBiz {
	return &userDisLikeRestaurantBiz{store: store, decStore: decStore}
}

func (biz userDisLikeRestaurantBiz) DisLikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDisLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.decStore.DecreaseLikeCount(ctx, restaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}