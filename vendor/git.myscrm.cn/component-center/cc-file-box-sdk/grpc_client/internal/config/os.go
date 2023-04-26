package config

import (
	"os"
)

const (
	// ElasticStack APM Server URL
	ENV_APM_SERVER_URL = "ELASTIC_APM_SERVER_URL"
	// ElasticStack APM Server URLs
	ENV_APM_SERVER_URLS = "ELASTIC_APM_SERVER_URLS"
	// Apollo META Server URL
	ENV_APOLLO_META_SERVER_URL = "APOLLO_META_SERVER_URL"
	// ETCD V3 Server URL
	ENV_ETCDV3_SERVER_URL = "ETCDV3_SERVER_URL"
	// ETCD V3 Server URLs
	ENV_ETCDV3_SERVER_URLS = "ETCDV3_SERVER_URLS"
	// TENANT_NAMESPACE
	TENANT_NAMESPACE = "TENANT_NAMESPACE"
	// GO_ENV
	GO_ENV = "GO_ENV"
	// YUNKE_ENV
	YUNKE_ENV = "YUNKE_ENV"
	// STARK_LOGGER_PATH
	STARK_LOGGER_PATH = "STARK_LOGGER_PATH"
	// APOLLO_ACCESSKEY_SECRET
	APOLLO_ACCESSKEY_SECRET = "APOLLO_ACCESSKEY_SECRET"
)

// GetAPMServerURLs gets elastic stack server url config from env.
func GetAPMServerURLs() string {
	value := os.Getenv(ENV_APM_SERVER_URL)
	if value != "" {
		return value
	}

	return os.Getenv(ENV_APM_SERVER_URLS)
}

// GetApolloServerURL gets apollo server url config from env.
func GetApolloServerURL() string {
	return os.Getenv(ENV_APOLLO_META_SERVER_URL)
}

// GetEtcdV3ServerURL gets etcd v3 server url config from env.
func GetEtcdV3ServerURL() string {
	return os.Getenv(ENV_ETCDV3_SERVER_URL)
}

// GetEtcdV3ServerURLs gets etcd v3 server urls config from env.
func GetEtcdV3ServerURLs() string {
	values := os.Getenv(ENV_ETCDV3_SERVER_URLS)
	if values != "" {
		return values
	}

	return GetEtcdV3ServerURL()
}

// GetTenantNamespace gets tenant namespace from env.
func GetTenantNamespace() string {
	return os.Getenv(TENANT_NAMESPACE)
}

// GetBuildEnv gets appEnv namespace from env.
func GetBuildEnv() string {
	value := os.Getenv(YUNKE_ENV)
	if value != "" {
		return value
	}

	value = os.Getenv(GO_ENV)
	if value != "" {
		return value
	}

	return ""
}

// GetStarkLoggerPath get logRootPath from env.
func GetStarkLoggerPath() string {
	return os.Getenv(STARK_LOGGER_PATH)
}

func GetApolloAccesskeySecret() string {
	return os.Getenv(APOLLO_ACCESSKEY_SECRET)
}
