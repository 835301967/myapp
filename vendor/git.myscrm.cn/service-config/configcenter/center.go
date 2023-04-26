package configcenter

import (
	"errors"
	"log"
	"strings"

	"github.com/tidwall/gjson"

	"git.myscrm.cn/golang/common/ykenv"
	"git.myscrm.cn/service-config/configcenter/read"
)

type ConfigCenter struct {
	parse      gjson.Result
	serverName string
	configName string
}

const (
	DEFAULT_GLOBAL_PATH      = "/usr/local/etc/global-conf"
	DEFAULT_GLOBAL_FILE_NAME = "/usr/local/etc/global-conf/config.json"
	DEFAULT_CONFIG_ROOT_PATH = "/usr/local/etc"

	DEFAULT_CERT_PEM_PATH      = "certs/server.pem"
	DEFAULT_CERT_KEY_PATH      = "certs/server-key.pem"
	DEFAULT_CERT_PEM_FILE_NAME = "server.pem"
	DEFAULT_CERT_KEY_FILE_NAME = "server-key.pem"

	SERVER_NAME      = "server_name"
	SERVER_PORT      = "server_port"
	CONFIG_ROOT_PATH = "config_root_path"
	CERT_PEM_PATH    = "cert_pem_path"
	CERT_KEY_PATH    = "cert_key_path"
	CERT_SERVER_NAME = "cert_server_name"
)

func NewConfigCenter(serverName string) *ConfigCenter {
	fileRead := read.NewFileRead()
	jsonByte, err := fileRead.Read(DEFAULT_GLOBAL_FILE_NAME)
	if err != nil {
		log.Fatalf("fileRead.Read err: %v", err)
	}

	env, err := ykenv.GetMode()
	if err != nil {
		log.Fatalf("ykenv.GetMode err: %v", err)
	}

	return &ConfigCenter{
		parse:      gjson.Parse(string(jsonByte[:])),
		serverName: serverName,
		configName: env + "/" + serverName,
	}
}

// 获取服务名称
func (c *ConfigCenter) GetServerName() (string, error) {
	value := c.parse.Get(c.getPath(SERVER_NAME))
	if !value.Exists() {
		return ``, errors.New(c.serverName + " server_name not exist")
	}

	return value.String(), nil
}

// 获取服务端口
func (c *ConfigCenter) GetServerPort() (uint64, error) {
	value := c.parse.Get(c.getPath(SERVER_PORT))
	if !value.Exists() {
		return 0, errors.New(c.serverName + " server_port not exist")
	}

	return value.Uint(), nil
}

// 获取服务配置目录
func (c *ConfigCenter) GetServerConfigPath() (string, error) {
	value := c.parse.Get(c.getPath(CONFIG_ROOT_PATH))
	if !value.Exists() {
		return ``, errors.New(c.serverName + " config_root_path not exist")
	}

	return strings.Join([]string{value.String(), c.configName}, "/"), nil
}

func (c *ConfigCenter) MustGetServerConfigPath() string {
	var paths []string
	value := c.parse.Get(c.getPath(CONFIG_ROOT_PATH))

	if !value.Exists() {
		paths = []string{DEFAULT_CONFIG_ROOT_PATH, c.configName}
	} else {
		paths = []string{value.String(), c.configName}
	}

	return strings.Join(paths, "/")
}

// 获取证书服务名称
func (c *ConfigCenter) GetCertServerName() (string, error) {
	value := c.parse.Get(c.getPath(CERT_SERVER_NAME))
	if !value.Exists() {
		return "", errors.New(c.serverName + " cert_server_name not exist")
	}

	return value.String(), nil
}

// 获取证书 .pem 路径
func (c *ConfigCenter) GetCertPemPath() (string, error) {
	value := c.parse.Get(c.getPath(CERT_PEM_PATH))
	if !value.Exists() {
		return ``, errors.New(c.serverName + " cert_pem_path not exist")
	}

	return value.String(), nil
}

func (c *ConfigCenter) MustGetCertPemPath() string {
	value := c.parse.Get(c.getPath(CERT_PEM_PATH))
	if !value.Exists() {
		return DEFAULT_CERT_PEM_PATH
	}

	return value.String()
}

// 获取证书 .key 路径
func (c *ConfigCenter) GetCertKeyPath() (string, error) {
	value := c.parse.Get(c.getPath(CERT_KEY_PATH))
	if !value.Exists() {
		return ``, errors.New(c.serverName + " cert_key_path not exist")
	} else {
		return value.String(), nil
	}
}

func (c *ConfigCenter) MustGetCertKeyPath() string {
	value := c.parse.Get(c.getPath(CERT_KEY_PATH))
	if !value.Exists() {
		return DEFAULT_CERT_KEY_PATH
	}

	return value.String()
}

func (c *ConfigCenter) getPath(key string) string {
	return strings.Join([]string{c.serverName, key}, ".")
}
