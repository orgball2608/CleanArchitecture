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

//type IncreaseLikeCountStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//incStore IncreaseLikeCountStore
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

	//// new side effect
	//job := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
	//	log.Println(err)
	//}
	return nil
}
