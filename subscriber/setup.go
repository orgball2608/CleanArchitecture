package subscriber

import (
	"LearnGo/component/appctx"
	"context"
)

func Setup(appCtx appctx.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
	DecreaseLikeCountAfterUserLikeRestaurant(
		appCtx, context.Background())
	PushNotificationAfterUserLikeRestaurant(
		appCtx, context.Background())
}
