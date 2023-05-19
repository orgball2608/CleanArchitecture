package subscriber

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantstorage "LearnGo/module/restaurant/storage"
	"context"
	"log"
)

func DecreaseLikeCountAfterUserLikeRestaurant(
	appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserDisLikeRestaurant)
	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			log.Println("Decrease like count")
			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
