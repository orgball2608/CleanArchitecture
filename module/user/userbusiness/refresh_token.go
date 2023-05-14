package userbusiness

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/component/tokenprovider"
	"LearnGo/module/user/usermodel"
	"context"
)

type RefreshStorage interface {
}

type refreshBusiness struct {
	appCtx             appctx.AppContext
	userStore          LoginStorage
	refreshTokenExpiry int // expiry will replace for type TokenConfig
	tokenProvider      tokenprovider.Provider
	hasher             Hasher
}

func NewRefreshBusiness(appCtx appctx.AppContext,
	userStore LoginStorage,
	refreshTokenExpiry int,
	tokenProvider tokenprovider.Provider,
	hasher Hasher) *refreshBusiness {
	return &refreshBusiness{
		appCtx:             appCtx,
		userStore:          userStore,
		refreshTokenExpiry: refreshTokenExpiry,
		tokenProvider:      tokenProvider,
		hasher:             hasher,
	}
}

// 1. Find user, email
// 2. Hash pass from input & compare with pass in db
// 3. Provider: issue JWT token for Client
// 3.1 Access token & Refresh token
// 4. Return token(s)

func (biz *refreshBusiness) Refresh(ctx context.Context, user interface{}) (*usermodel.RefreshTokenResponse, error) {

	data, ok := user.(*usermodel.User)

	if !ok {
		return nil, common.ErrInternal(nil)
	}

	payload := tokenprovider.TokenPayload{
		UserId: data.GetUserId(),
		Role:   data.GetUserRole(),
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.refreshTokenExpiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewRefreshTokenResponse(refreshToken)

	return account, nil
}
