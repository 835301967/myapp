package client_conn

import (
	"context"
	"errors"
	"fmt"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/internal/config"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/internal/service"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/internal/service/etcdconfig"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/module/grpc_interceptor"
	"os"
	"strings"

	"git.myscrm.cn/golang/elastic-apm/module/apmgrpc"
	"git.myscrm.cn/service-config/configcenter"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

type Conn struct {
	ServerName     string
	ServerPort     string
	CertFile       string
	CertServerName string
	IsSsl          bool

	target func(ctx context.Context) (string, error)

	chainUnaryInts  []grpc.UnaryClientInterceptor
	chainStreamInts []grpc.StreamClientInterceptor
}

func NewConn(serviceName string) (*Conn, error) {
	center := configcenter.NewConfigCenterV2(serviceName)
	certFile, err := center.GetCertPemPath()
	if err != nil {
		return nil, fmt.Errorf("NewConn.center.GetCertPemPath err: %v", err)
	}

	etcdServerUrls := config.GetEtcdV3ServerURLs()
	if len(etcdServerUrls) == 0 {
		return nil, fmt.Errorf("Can't not found env '%s'", config.ENV_ETCDV3_SERVER_URLS)
	}
	serviceLB := service.NewService(etcdServerUrls, serviceName)
	serviceConfig := etcdconfig.NewServiceConfig(serviceLB)
	currentConfig, _, err := serviceConfig.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("serviceConfig.GetConfig err: %v", err)
	}

	serviceNames := strings.Split(serviceName, "-")
	if len(serviceNames) < 1 {
		return nil, errors.New("NewConn.serviceNames is empty")
	}

	certServerName := ""
	exists, _ := IsFileExists(certFile)
	if exists == true {
		certServerName = serviceNames[0]
	} else {
		certFile = ""
	}

	return &Conn{
		ServerName:     serviceName,
		ServerPort:     currentConfig.ServicePort,
		CertFile:       certFile,
		CertServerName: certServerName,
		IsSsl:          currentConfig.IsSsl,
		target: func(ctx context.Context) (string, error) {
			return serviceName + ":" + currentConfig.ServicePort, nil
		},
	}, nil
}

func IsFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (c *Conn) WithTarget(fn func(context.Context) (string, error)) *Conn {
	c.target = fn
	return c
}

func (c *Conn) WithChainUnaryInterceptor(interceptors ...grpc.UnaryClientInterceptor) *Conn {
	c.chainUnaryInts = append(c.chainUnaryInts, interceptors...)
	return c
}

func (c *Conn) WithStreamClientInterceptor(interceptors ...grpc.StreamClientInterceptor) *Conn {
	c.chainStreamInts = append(c.chainStreamInts, interceptors...)
	return c
}

// Get ElasticStack gRPC APM Conn(v2)
func (c *Conn) GetAPMConn(ctx context.Context, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if c.IsSsl {
		creds, err := credentials.NewClientTLSFromFile(c.CertFile, c.CertServerName)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 一元拦截器注册
	chainUnaryClients := []grpc.UnaryClientInterceptor{
		apmgrpc.NewUnaryClientInterceptor(),
		grpc_interceptor.UnaryCtxGRPC(),
		grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(2),
			grpc_retry.WithCodes(
				codes.Internal,
				codes.DeadlineExceeded,
			),
		),
	}
	if len(c.chainUnaryInts) > 0 {
		chainUnaryClients = append(chainUnaryClients, c.chainUnaryInts...)
	}
	opts = append(opts, grpc.WithChainUnaryInterceptor(chainUnaryClients...))

	target, err := c.target(ctx)
	if err != nil {
		return nil, err
	}

	return grpc.DialContext(
		ctx,
		target,
		opts...,
	)
}
