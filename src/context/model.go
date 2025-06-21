package context

import core "github.com/cynxees/cynx-core/proto/gen"

type Key string

const (
	KeyRequestId     Key = "request_id"
	KeyRequestOrigin Key = "request_origin"
	KeyRequestPath   Key = "request_path"

	KeyUsername Key = "username"
	KeyUserId   Key = "user_id"   // int32
	KeyUserType Key = "user_type" // int32

	KeyBaseRequest Key = "base_request" // *pb.BaseRequest (protobuf message for base request info
)

type RequestWithBase interface {
	GetBase() *core.BaseRequest
}
