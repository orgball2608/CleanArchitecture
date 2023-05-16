package ginrestaurant

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantbiz "LearnGo/module/restaurant/biz"
	restaurantmodel "LearnGo/module/restaurant/model"
	restaurantstorage "LearnGo/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		data.UserId = requester.GetUserId()

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
