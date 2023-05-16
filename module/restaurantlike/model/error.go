package restaurantlikemodel

import (
	"LearnGo/common"
	"fmt"
)

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"))
}
func ErrCannotDisLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("cannot unlike this restaurant"),
		fmt.Sprintf("ErrCannotUnLikeRestaurant"))
}
