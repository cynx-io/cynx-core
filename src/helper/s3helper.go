package helper

import (
	"context"
	"time"

	"github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/externalapi/s3"
	"google.golang.org/grpc/codes"
)

func HandleUploadFile(ctx context.Context, req *core.UploadFileRequest, resp *core.UploadFileResponse) error {
	if req.Base == nil {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Base request is required",
		}
		return nil
	}

	if req.Bucket == "" {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Bucket is required",
		}
		return nil
	}

	if req.Key == "" {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Key is required",
		}
		return nil
	}

	if req.ContentType == "" {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Content type is required",
		}
		return nil
	}

	if len(req.FileData) == 0 {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "File data is required",
		}
		return nil
	}

	result, err := s3.UploadFile(ctx, req.Bucket, req.Key, req.ContentType, req.FileData)
	if err != nil {
		resp.Base = &core.BaseResponse{
			Code: codes.Internal.String(),
			Desc: "Failed to upload file: " + err.Error(),
		}
		return nil
	}

	resp.Base = &core.BaseResponse{
		Code: codes.OK.String(),
		Desc: "File uploaded successfully",
	}
	resp.Bucket = result.Bucket
	resp.Key = result.Key
	resp.Location = result.Location
	resp.Etag = result.ETag

	return nil
}

func HandleGeneratePresignedURL(ctx context.Context, req *core.GeneratePresignedURLRequest, resp *core.GeneratePresignedURLResponse) error {
	if req.Base == nil {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Base request is required",
		}
		return nil
	}

	if req.Bucket == "" {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Bucket is required",
		}
		return nil
	}

	if req.Key == "" {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Key is required",
		}
		return nil
	}

	if req.ContentType == "" {
		resp.Base = &core.BaseResponse{
			Code: codes.InvalidArgument.String(),
			Desc: "Content type is required",
		}
		return nil
	}

	expiresIn := time.Duration(req.ExpiresInSeconds) * time.Second
	if expiresIn <= 0 {
		expiresIn = 15 * time.Minute // Default to 15 minutes
	}

	url, err := s3.GeneratePresignedUploadURL(ctx, req.Bucket, req.Key, req.ContentType, expiresIn)
	if err != nil {
		resp.Base = &core.BaseResponse{
			Code: codes.Internal.String(),
			Desc: "Failed to generate presigned URL: " + err.Error(),
		}
		return nil
	}

	resp.Base = &core.BaseResponse{
		Code: codes.OK.String(),
		Desc: "Presigned URL generated successfully",
	}
	resp.UploadUrl = url

	return nil
}
