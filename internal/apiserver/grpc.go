package apiserver

import (
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"goer-startup/internal/apiserver/controller/v1/user"
	"goer-startup/internal/apiserver/store"
	"goer-startup/internal/pkg/log"
	pb "goer-startup/pkg/proto/goer/v1"
)

// startGRPCServer 创建并运行 GRPC 服务器.
func startGRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
	if err != nil {
		log.Fatalw("Failed to listen", "err", err)
	}

	// 创建 GRPC Server 实例
	grpcSrv := grpc.NewServer()
	pb.RegisterGoerServer(grpcSrv, user.NewUserController(store.S, nil))

	// 运行 GRPC 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 GRPC 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on grpc address", "addr", viper.GetString("grpc.addr"))
	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalw(err.Error())
		}
	}()

	return grpcSrv
}
