package ginrestaurant

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantbiz "LearnGo/module/restaurant/biz"
	restaurantstorage "LearnGo/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			// Any err thrown from Biz belongs to Application error
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
