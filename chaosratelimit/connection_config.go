package chaosratelimit

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/schema"
)

type chaosratelimiterConfig struct {
	Regions []string `cty:"regions"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"regions": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
}

func ConfigInstance() interface{} {
	return &chaosratelimiterConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) chaosratelimiterConfig {
	if connection == nil {
		return chaosratelimiterConfig{}
	}
	raw := connection.GetConfig()
	if raw == nil {
		return chaosratelimiterConfig{}
	}
	config, _ := raw.(chaosratelimiterConfig)
	return config
}
