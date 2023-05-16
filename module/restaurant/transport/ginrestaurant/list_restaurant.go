package ginrestaurant

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	restaurantbiz "LearnGo/module/restaurant/biz"
	restaurantmodel "LearnGo/module/restaurant/model"
	restaurantrepo "LearnGo/module/restaurant/repository"
	restaurantstorage "LearnGo/module/restaurant/storage"
	restaurantlikestorage "LearnGo/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantBiz(store, likeStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)
		pagingData.Fulfill()

		filter.Status = 1

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
