package client

import (
	"LearnGo/component/appctx"
	"LearnGo/middleware"
	"LearnGo/module/restaurant/transport/ginrestaurant"
	ginrestaurantlike "LearnGo/module/restaurantlike/transport/ginrstlike"
	"github.com/gin-gonic/gin"
)

func ClientRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	restaurants := v1.Group("restaurants", middleware.RequireAuth(appContext))

	// POST /restaurants
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	// GET all restaurants
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	// GET restaurants by ID
	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appContext))

	// DELETE Restaurant by id
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	// UPDATE Restaurant by id
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))

	//CREATE Restaurant like
	restaurants.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appContext))

	//DELETE Restaurant Dislike
	restaurants.DELETE("/:id/dislike", ginrestaurantlike.UserDisLikeRestaurant(appContext))

	restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUsersLikeRestaurant(appContext))
}
