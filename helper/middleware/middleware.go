package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func ValidateUserInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Thực hiện kiểm tra tính hợp lệ ở đây
		// Nếu có vấn đề, trả về lỗi
		// Nếu không, tiếp tục xử lý request
		fmt.Println("Middleware: Kiểm tra tính hợp lệ của tên người dùng")
		return handler(ctx, req)
	}
}
