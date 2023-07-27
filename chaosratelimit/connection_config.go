package chaosratelimit

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
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
	if connection == nil || connection.Config == nil {
		return chaosratelimiterConfig{}
	}
	config, _ := connection.Config.(chaosratelimiterConfig)
	return config
}
