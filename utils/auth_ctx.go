package utils

import (
	"context"
	"lms_system/internal/domain/common"
	"lms_system/internal/domain/entity"
)

func GetUserFromContext(ctx context.Context) *entity.UserContext {
	userCtx, ok := ctx.Value(common.UserContextKey).(*entity.UserContext)
	if !ok {
		return nil
	}
	return userCtx
}
