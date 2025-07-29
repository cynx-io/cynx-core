package grpc

import (
	"context"
	"encoding/json"
	core "github.com/cynx-io/cynx-core/proto/gen"
	coreContext "github.com/cynx-io/cynx-core/src/context"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/cynx-core/src/response"
	"reflect"
	"time"
)

type RequestWithBase interface {
	GetBase() *core.BaseRequest
}

func HandleGrpc[Req RequestWithBase, Resp response.Generic](
	ctx context.Context,
	req Req,
	_ Resp,
	serviceFunc func(context.Context, Req, Resp) error,
) (Resp, error) {

	ctx, err := coreContext.SetBaseRequest(ctx, req.GetBase())
	if err != nil {
		logger.Error(ctx, "Failed to set context base request: ", err)
		// Continue either way
	}

	localReq := req
	go func() {
		reqBase := localReq.GetBase()
		rawBody, _ := json.Marshal(req)
		ctxBg := context.Background()
		err := logger.LogTrxElasticsearch(ctxBg, logger.TrxEntry{
			Timestamp:     time.Now(),
			UserId:        reqBase.UserId,
			Username:      reqBase.Username,
			UserType:      reqBase.UserType,
			IpAddress:     reqBase.IpAddress,
			RequestId:     reqBase.RequestId,
			RequestOrigin: reqBase.RequestOrigin,
			Endpoint:      reqBase.RequestPath,
			Type:          "REQUEST",
			Body:          rawBody,
		})
		if err != nil {
			logger.Error(ctxBg, "Failed to log transaction to Elasticsearch: ", err)
		}
	}()

	resp := newResponse[Resp]()
	baseResp := setProtoBaseResponse(resp)

	if v, ok := any(req).(interface{ ValidateAll() error }); ok {
		if err := v.ValidateAll(); err != nil {

			baseResp.Code = "VE"
			baseResp.Desc += ": " + err.Error()

			return resp, nil
		}
	}

	err = serviceFunc(ctx, req, resp)
	if err != nil {
		baseResp.Desc += ": " + err.Error()
	}

	localResp := resp
	go func() {
		reqBase := localReq.GetBase()
		rawBody, _ := json.Marshal(localResp)
		ctxBg := context.Background()
		err := logger.LogTrxElasticsearch(ctxBg, logger.TrxEntry{
			Timestamp:     time.Now(),
			UserId:        reqBase.UserId,
			Username:      reqBase.Username,
			UserType:      reqBase.UserType,
			IpAddress:     reqBase.IpAddress,
			RequestId:     reqBase.RequestId,
			RequestOrigin: reqBase.RequestOrigin,
			Endpoint:      reqBase.RequestPath,
			Type:          "RESPONSE",
			Body:          rawBody,
		})
		if err != nil {
			logger.Error(ctxBg, "Failed to log transaction to Elasticsearch: ", err)
		}
	}()

	return resp, nil
}

func setProtoBaseResponse(resp any) *core.BaseResponse {
	v := reflect.ValueOf(resp)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	base := &core.BaseResponse{}
	field := v.FieldByName("Base")
	if field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(base))
	}

	return base
}

func newResponse[T any]() T {
	var ptr T
	ptrType := reflect.TypeOf(ptr)

	if ptrType.Kind() == reflect.Ptr {
		elem := reflect.New(ptrType.Elem())
		return elem.Interface().(T)
	}

	return ptr
}
