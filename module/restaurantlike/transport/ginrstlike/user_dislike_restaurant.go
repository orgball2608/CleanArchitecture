package ginrestaurantlike

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantlikebussiness "LearnGo/module/restaurantlike/bussiness"
	restaurantlikestorage "LearnGo/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserDisLikeRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(db)
		//decStore := restaurantstorage.NewSQLStore(db)
		biz := restaurantlikebussiness.NewUserDisLikeRestaurantBiz(store, ctx.GetPubSub())

		if err := biz.DisLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(id.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse("true"))

	}
}
