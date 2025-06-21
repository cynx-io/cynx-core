package grpc

import (
	"context"
	core "github.com/cynxees/cynx-core/proto/gen"
	"github.com/cynxees/cynx-core/src/response"
	"reflect"
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

	resp := newResponse[Resp]()
	base := setProtoBase(resp)

	if v, ok := any(req).(interface{ ValidateAll() error }); ok {
		if err := v.ValidateAll(); err != nil {

			base.Code = "VE"
			base.Desc += ": " + err.Error()

			return resp, nil
		}
	}

	err := serviceFunc(ctx, req, resp)

	if err != nil {
		base.Desc += ": " + err.Error()
	}

	return resp, nil
}

func setProtoBase(resp any) *core.BaseResponse {
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
