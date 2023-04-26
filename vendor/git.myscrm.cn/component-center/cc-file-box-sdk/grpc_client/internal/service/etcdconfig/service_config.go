package etcdconfig

import (
	"context"
	"fmt"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/common/json"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/internal/service"
	"git.myscrm.cn/component-center/cc-file-box-sdk/grpc_client/internal/util"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"go.etcd.io/etcd/client"
)

const (
	ROOT            = "/"
	SERVICE         = "service"
	DEFALUT_CLUSTER = "default"
)

type ServiceConfig struct {
	ServiceLB *service.ServiceLB
	Config
}

type Config struct {
	ServiceVersion string `json:"service_version"`
	ServicePort    string `json:"service_port"`
	HttpPort       string `json:"http_port"`
	IsSsl          bool   `json:"is_ssl"`
}

func NewServiceConfig(slb *service.ServiceLB) *ServiceConfig {
	return &ServiceConfig{ServiceLB: slb}
}

func (s *ServiceConfig) GetKeyName(serverName string) string {
	return ROOT + SERVICE + "." + serverName + "." + DEFALUT_CLUSTER
}

func (s *ServiceConfig) GetConfig() (*Config, bool, error) {
	cli, err := util.NewEtcd(s.ServiceLB.EtcdServerUrl)
	if err != nil {
		return nil, false, fmt.Errorf("util.NewEtcdKeysAPI err: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := s.GetKeyName(s.ServiceLB.ServerName)
	serviceInfo, err := client.NewKeysAPI(cli).Get(ctx, key, nil)
	if err != nil {
		return nil, client.IsKeyNotFound(err), fmt.Errorf("cli.Get err: %v", err)
	}

	var config Config
	if len(serviceInfo.Node.Value) > 0 {
		err = json.Unmarshal(serviceInfo.Node.Value, &config)
		if err != nil {
			return nil, true, fmt.Errorf("json.Unmarshal err: %v", err)
		}

		// 修复：v2.3.3 起使用 is_ssl 作为判断依据，但先前的版本中 is_ssl 是没有强制写入的
		// 因此会造成 is_ssl 不存在，导致以非TLS的模式请求，但18-19年的核心版本是TLS Server。
		if !gjson.Get(serviceInfo.Node.Value, "is_ssl").Exists() {
			config.IsSsl = true
		}
	}

	if config.ServicePort == "" {
		return nil, true, fmt.Errorf("servicePort is empty, key: %s", key)
	}

	return &config, false, nil
}

func (s *ServiceConfig) WriteConfig(c Config) error {
	cli, err := util.NewEtcd(s.ServiceLB.EtcdServerUrl)
	if err != nil {
		return fmt.Errorf("util.NewEtcdKeysAPI err: %v", err)
	}

	key := s.GetKeyName(s.ServiceLB.ServerName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sconfig, err := json.MarshalToString(&c)
	if err != nil {
		return fmt.Errorf("json.MarshalToString err: %v", err)
	}

	_, err = client.NewKeysAPI(cli).Set(ctx, key, sconfig, nil)
	if err != nil {
		return fmt.Errorf("cli.Put err: %v", err)
	}

	return nil
}

func (s *ServiceConfig) GetConfigs() (map[string]*Config, error) {
	cli, err := util.NewEtcd(s.ServiceLB.EtcdServerUrl)
	if err != nil {
		return nil, fmt.Errorf("util.NewEtcdKeysAPI err: %v", err)
	}

	kapi := client.NewKeysAPI(cli)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serviceInfos, err := kapi.Get(ctx, "/", nil)
	if err != nil {
		return nil, fmt.Errorf("cli.Get err: %v", err)
	}

	configs := make(map[string]*Config)
	for _, info := range serviceInfos.Node.Nodes {
		if len(info.Value) > 0 {
			index := strings.Index(info.Key, ROOT+SERVICE)
			if index == 0 {
				config := &Config{}
				err := json.Unmarshal(info.Value, config)
				if err != nil {
					return nil, fmt.Errorf("json.UnmarshalByte err: %v", err)
				}

				configs[info.Key] = config
			}
		}
	}

	return configs, nil
}
