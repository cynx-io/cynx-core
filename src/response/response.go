package response

import core "github.com/cynxees/cynx-core/proto/gen"

type Generic interface {
	GetBase() *core.BaseResponse
}
