package conn

import (
	"contentService/pkg/config"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"strconv"
)

func NewRpcConnClient(conf config.GConfig) (*grpc.ClientConn, func(), error) {
	zConf := zrpc.RpcClientConf{
		Target: conf.Rpc.ListenIP + ":" + strconv.Itoa(conf.Rpc.RPCPort),
	}
	client, err := zrpc.NewClient(zConf, zrpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return nil, func() {}, err
	}
	conn := client.Conn()
	return conn, func() {
		_ = conn.Close()
	}, nil
}
