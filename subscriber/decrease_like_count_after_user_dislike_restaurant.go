package subscriber

import (
	"LearnGo/component/appctx"
	restaurantstorage "LearnGo/module/restaurant/storage"
	"LearnGo/pubsub"
	"context"
)

func DecreaseLikeCountAfterUserLikeRestaurant(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
