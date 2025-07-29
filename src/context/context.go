package context

import (
	"context"
	core "github.com/cynx-io/cynx-core/proto/gen"
)

func SetupContext(ctx context.Context, req RequestWithBase) context.Context {
	if req == nil {
		return ctx
	}

	baseReq := req.GetBase()
	if baseReq == nil {
		return ctx
	}

	ctx, err := SetBaseRequest(ctx, baseReq)
	if err != nil {
		return ctx
	}

	if baseReq.RequestId != "" {
		ctx = SetKey(ctx, KeyRequestId, baseReq.RequestId)
	}

	if baseReq.RequestOrigin != "" {
		ctx = SetKey(ctx, KeyRequestOrigin, baseReq.RequestOrigin)
	}

	if baseReq.RequestPath != "" {
		ctx = SetKey(ctx, KeyRequestPath, baseReq.RequestPath)
	}

	if baseReq.Username != nil {
		ctx = SetKey(ctx, KeyUsername, *baseReq.Username)
	}

	if baseReq.UserId != nil {
		ctx = SetUserId(ctx, *baseReq.UserId)
	}

	if baseReq.UserType != nil {
		ctx = SetUserType(ctx, *baseReq.UserType)
	}

	return ctx
}

func GetBaseRequest(ctx context.Context) *core.BaseRequest {
	val := ctx.Value(KeyBaseRequest)
	if val == nil {
		return nil
	}
	if req, ok := val.(*core.BaseRequest); ok {
		return req
	}
	return nil
}

func SetBaseRequest(ctx context.Context, req *core.BaseRequest) (context.Context, error) {
	return context.WithValue(ctx, KeyBaseRequest, req), nil
}

func SetKey(ctx context.Context, key Key, value string) context.Context {
	if value == "" {
		return ctx
	}
	return context.WithValue(ctx, key, value)
}

func SetUserId(ctx context.Context, userId int32) context.Context {
	return context.WithValue(ctx, KeyUserId, userId)
}

func SetUserType(ctx context.Context, userType int32) context.Context {
	return context.WithValue(ctx, KeyUserType, userType)
}

func GetKeyOrEmpty(ctx context.Context, key Key) string {
	value := ctx.Value(key)
	if value == nil {
		return ""
	}

	strValue, ok := value.(string)
	if !ok {
		return ""
	}

	return strValue
}

func GetKey(ctx context.Context, key Key) *string {
	value := ctx.Value(key)
	if value == nil {
		return nil
	}

	strValue, ok := value.(string)
	if !ok {
		return nil
	}

	return &strValue
}

func GetUserId(ctx context.Context) *int32 {
	value := ctx.Value(KeyUserId)
	if value == nil {
		return nil
	}

	userId, ok := value.(int32)
	if !ok {
		return nil
	}
	return &userId
}

func GetUserType(ctx context.Context) *int32 {
	value := ctx.Value(KeyUserType)
	if value == nil {
		return nil
	}

	userType, ok := value.(int32)
	if !ok {
		return nil
	}
	return &userType
}
