package ginrestaurantlike

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantlikebussiness "LearnGo/module/restaurantlike/bussiness"
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	restaurantlikestorage "LearnGo/module/restaurantlike/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLikeRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data = restaurantlikemodel.UserLike{
			RestaurantId: int(id.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		fmt.Println(data)

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantlikestorage.NewSQLStore(db)

		biz := restaurantlikebussiness.NewUserLikeRestaurantBiz(store)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse("true"))

	}
}
