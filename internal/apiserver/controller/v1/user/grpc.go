package user

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"goer-startup/internal/pkg/log"
	pb "goer-startup/pkg/proto/goer/v1"
)

// ListUser 返回用户列表，只有 root 用户才能获取用户列表.
func (ctrl *UserController) ListUser(ctx context.Context, r *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	log.C(ctx).Infow("ListUser function called")

	resp, err := ctrl.b.Users().List(ctx, int(r.Offset), int(r.Limit))
	if err != nil {
		return nil, err
	}

	users := make([]*pb.UserInfo, 0, len(resp.Users))
	for _, u := range resp.Users {
		users = append(users, &pb.UserInfo{
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Phone,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
		})
	}

	ret := &pb.ListUserResponse{
		TotalCount: resp.TotalCount,
		Users:      users,
	}

	return ret, nil
}
