package subscriber

import (
	"LearnGo/component/appctx"
	restaurantstorage "LearnGo/module/restaurant/storage"
	"LearnGo/pubsub"
	"context"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotificationAfterUserLikeRestaurant(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push Notification like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push Notification when user like restaurant", likeData)
			return nil
		},
	}
}
