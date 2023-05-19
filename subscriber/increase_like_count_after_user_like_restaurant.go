package subscriber

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantstorage "LearnGo/module/restaurant/storage"
	"context"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(
	appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	log.Println("Increase like count")

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			log.Println("Increase like count")
			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}

func PushNotificationAfterUserLikeRestaurant(
	appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserDisLikeRestaurant)

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			log.Println("Push Notification when user like restaurant", likeData)
		}
	}()
}
