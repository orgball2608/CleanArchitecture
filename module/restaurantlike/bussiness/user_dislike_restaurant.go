package restaurantlikebussiness

import (
	"LearnGo/common"
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	"LearnGo/pubsub"
	"context"
)

type UserDisLikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type userDisLikeRestaurantBiz struct {
	store  UserDisLikeRestaurantStore
	pubsub pubsub.Pubsub
}

func NewUserDisLikeRestaurantBiz(
	store UserDisLikeRestaurantStore, ps pubsub.Pubsub) *userDisLikeRestaurantBiz {
	return &userDisLikeRestaurantBiz{store: store, pubsub: ps}
}

func (biz userDisLikeRestaurantBiz) DisLikeRestaurant(
	ctx context.Context,
	userId,
	restaurantId int,
) error {

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikemodel.ErrCannotDisLikeRestaurant(err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserDisLikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{
		RestaurantId: restaurantId,
		UserId:       userId,
	}))
	return nil
}
