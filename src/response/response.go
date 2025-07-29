package response

import core "github.com/cynx-io/cynx-core/proto/gen"

type Generic interface {
	GetBase() *core.BaseResponse
}
