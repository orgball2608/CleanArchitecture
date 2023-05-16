package client

import (
	"LearnGo/component/appctx"
	"LearnGo/middleware"
	"LearnGo/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
)

func ClientRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	restaurants := v1.Group("restaurants", middleware.RequireAuth(appContext))

	// POST /restaurants

	restaurants.POST("", middleware.RequireAuth(appContext), ginrestaurant.CreateRestaurant(appContext))

	// GET all restaurants

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	// GET restaurants by ID
	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appContext))

	// DELETE Restaurant by id

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	// UPDATE Restaurant by id

	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))
}
