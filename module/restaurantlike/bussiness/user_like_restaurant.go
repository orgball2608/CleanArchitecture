package restaurantlikebussiness

import (
	"LearnGo/common"
	restaurantmodel "LearnGo/module/restaurant/model"
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"LearnGo/pubsub"
	"context"
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
	store  UserLikeRestaurantStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore, ps pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, pubsub: ps}
}

func (biz userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.UserLike) error {
	if err := biz.store.CreateLikeRestaurant(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))
	return nil
}
