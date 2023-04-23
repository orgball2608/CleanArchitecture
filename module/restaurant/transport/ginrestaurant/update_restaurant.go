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

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
