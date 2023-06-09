package ginrestaurantlike

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantlikebussiness "LearnGo/module/restaurantlike/bussiness"
	restaurantlikemodel "LearnGo/module/restaurantlike/model"
	restaurantlikestorage "LearnGo/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsersLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		//myArr := []string{}
		//
		//fmt.Println(myArr[0])

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebussiness.NewListUsersLikeRestaurantBiz(store)

		users, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range users {
			users[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))
	}
}
