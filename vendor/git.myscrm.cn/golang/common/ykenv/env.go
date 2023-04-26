package ykenv

import (
	"errors"
	"os"
	"strings"
)

const (
	ENV_NAME = "GO_ENV" //标识名

	DEV_MODE     = "dev"     //开发环境
	TEST_MODE    = "test"    //测试环境
	RELEASE_MODE = "release" //预发布环境
	PROD_MODE    = "prod"    //生产环境
)

// 获取当前环境的运行模式
func GetMode() (env string, err error) {
	env = strings.ToLower(os.Getenv(ENV_NAME))
	if env == "" {
		err = errors.New("Can not find ENV '" + ENV_NAME + "'")
	}

	return env, err
}

// 判断是否为开发模式
func IsDevMode() bool {
	return checkMode(DEV_MODE)
}

// 判断是否测试模式
func IsTestMode() bool {
	return checkMode(TEST_MODE)
}

// 判断是否预发布模式
func IsReleaseMode() bool {
	return checkMode(RELEASE_MODE)
}

// 判断是否生产环境
func IsProdMode() bool {
	return checkMode(PROD_MODE)
}

// 检查当前运行模式
func checkMode(mode string) bool {
	env := strings.ToLower(os.Getenv(ENV_NAME))
	return env == mode
}
