package restaurantlikebussiness

import (
	restaurantmodel "LearnGo/module/restaurant/model"
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"context"
	"fmt"
)

type UserLikeRestaurantStore interface {
	CreateLikeRestaurant(context context.Context, data *restaurantlikemodel.UserLike) error
}

type IncreaseLikeCount interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store}
}

func (biz userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.UserLike) error {
	fmt.Printf("data", *data)

	if err := biz.store.CreateLikeRestaurant(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
