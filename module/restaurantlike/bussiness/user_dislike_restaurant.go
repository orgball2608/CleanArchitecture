package restaurantlikebussiness

import (
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"context"
)

type UserDisLikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type userDisLikeRestaurantBiz struct {
	store UserDisLikeRestaurantStore
}

func NewUserDisLikeRestaurantBiz(
	store UserDisLikeRestaurantStore) *userDisLikeRestaurantBiz {
	return &userDisLikeRestaurantBiz{store: store}
}

func (biz userDisLikeRestaurantBiz) DisLikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDisLikeRestaurant(err)
	}

	return nil
}
