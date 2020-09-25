/**
 * Copyright 2015 @ at3.net.
 * name : client.go
 * author : jarryliu
 * date : 2016-11-13 12:32
 * description :
 * history :
 */
package service

import (
	"github.com/ixre/gof/log"
	"go.etcd.io/etcd/clientv3"
	"go2o/core/etcd"
	"go2o/core/service/proto"
	"google.golang.org/grpc"
	"os"
	"time"
)

//var cfg  clientv3.Config
var selector etcd.Selector

// 设置Thrift地址
func ConfigureClient(c clientv3.Config) {
	//cfg = c
	log.Println("[ Go2o][ RPC]: connecting go2o rpc server...")
	s, err := etcd.NewSelector(service, c, etcd.AlgRoundRobin)
	if err != nil {
		log.Println("[ Go2o][ RPC]: can't connect go2o rpc server! ", err.Error())
		os.Exit(1)
	}
	selector = s
	go tryConnect(10)
}

// 尝试连接服务,如果连接不成功,则退出
func tryConnect(retryTimes int) {
	time.Sleep(time.Second)
	for i := 0; i < retryTimes; i++ {
		trans, _, err := StatusServiceClient()
		if err == nil {
			trans.Close()
			break
		} else if i == retryTimes-1 {
			log.Println("[ Go2o][ RPC]: can't connect go2o rpc server! ", err.Error())
			os.Exit(1)
		}
		time.Sleep(time.Second * 2)
	}
}

// 获取连接
func getConn(selector etcd.Selector) (*grpc.ClientConn, error) {
	next, err := selector.Next()
	if err != nil {
		log.Printf("[ go2o][ error]: %s\n", err.Error())
		return nil, err
	}
	return grpc.Dial(next.Addr, grpc.WithInsecure(), grpc.WithBlock())
}

// 状态客户端
func StatusServiceClient() (*grpc.ClientConn, proto.StatusServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewStatusServiceClient(conn), err
	}
	return conn, nil, err
}

// 基础服务
func RegistryServiceClient() (*grpc.ClientConn, proto.RegistryServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewRegistryServiceClient(conn), err
	}
	return conn, nil, err
}

// 商户客户端
func MerchantServiceClient() (*grpc.ClientConn, proto.MerchantServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewMerchantServiceClient(conn), err
	}
	return conn, nil, err
}

// 会员客户端
func MemberServiceClient() (*grpc.ClientConn, proto.MemberServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewMemberServiceClient(conn), err
	}
	return conn, nil, err
}

// 基础服务
func FoundationServiceClient() (*grpc.ClientConn, proto.FoundationServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewFoundationServiceClient(conn), err
	}
	return conn, nil, err
}

// 消息客户端
func MessageServiceClient() (*grpc.ClientConn, proto.MessageServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewMessageServiceClient(conn), err
	}
	return conn, nil, err
}

// 消息客户端
func ContentServiceClient() (*grpc.ClientConn, proto.ContentServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewContentServiceClient(conn), err
	}
	return conn, nil, err
}

// 支付服务
func PaymentServiceClient() (*grpc.ClientConn, proto.PaymentServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewPaymentServiceClient(conn), err
	}
	return conn, nil, err
}

// 钱包服务
func WalletClient() (*grpc.ClientConn, proto.WalletServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewWalletServiceClient(conn), err
	}
	return conn, nil, err
}

// 订单服务
func OrderServiceClient() (*grpc.ClientConn, proto.OrderServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewOrderServiceClient(conn), err
	}
	return conn, nil, err
}

// 快递服务
func ExpressServiceClient() (*grpc.ClientConn, proto.ExpressServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewExpressServiceClient(conn), err
	}
	return conn, nil, err
}

// 物流服务
func ShipmentServiceClient() (*grpc.ClientConn, proto.ShipmentServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewShipmentServiceClient(conn), err
	}
	return conn, nil, err
}

// 商品服务
func ItemServiceClient() (*grpc.ClientConn, proto.ItemServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewItemServiceClient(conn), err
	}
	return conn, nil, err
}

// 产品服务
func ProductServiceClient() (*grpc.ClientConn, proto.ProductServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewProductServiceClient(conn), err
	}
	return conn, nil, err
}

// 商店服务
func ShopServiceClient() (*grpc.ClientConn, proto.ShopServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewShopServiceClient(conn), err
	}
	return conn, nil, err
}

// 财务服务
func FinanceServiceClient() (*grpc.ClientConn, proto.FinanceServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewFinanceServiceClient(conn), err
	}
	return conn, nil, err
}

// 查询服务
func QueryServiceClient() (*grpc.ClientConn, proto.QueryServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewQueryServiceClient(conn), err
	}
	return conn, nil, err
}

// 售后服务
func AfterSalesServiceClient() (*grpc.ClientConn, proto.AfterSalesServiceClient, error) {
	conn, err := getConn(selector)
	if err == nil {
		return conn, proto.NewAfterSalesServiceClient(conn), err
	}
	return conn, nil, err
}
