package chaosratelimit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/rate_limiter"
)

const (
	pluginName              = "steampipe-provider-chaosratelimit"
	rateLimiterScopeService = "service"
	rateLimiterScopeRegion  = "region"
)

func Plugin(_ context.Context) *plugin.Plugin {
	//time.Sleep(10 * time.Second)
	p := &plugin.Plugin{
		Name: pluginName,
		TableMap: map[string]*plugin.Table{
			"chaosratelimit_rate_limiter":      rateLimiterTable(),
			"chaosratelimit_list_parent_child": parentChildTable()},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		RateLimiters: []*rate_limiter.Definition{
			{
				Name:       "limiter1",
				FillRate:   5,
				BucketSize: 1,
				Scope:      []string{"limiter_rate"},
				Where:      "limiter_rate = '3'",
			},
			//{
			//	Name:       "limiter_rate scoped limiter, 2/s",
			//	FillRate:   2,
			//	BucketSize: 1,
			//	Scopes:     []string{"limiter_rate"},
			//	Where:      "limiter_rate = '2'",
			//},
			//{
			//	Name:       "limiter_rate scoped limiter, 1/s",
			//	FillRate:   1,
			//	BucketSize: 1,
			//	Scopes:     []string{"limiter_rate"},
			//	Where:      "limiter_rate = '1'",
			//},
			//{
			//	Name:       "region scoped limiter (column scope)",
			//	FillRate:   10,
			//	BucketSize: 10,
			//	Scopes:     []string{rateLimiterScopeRegion},
			//},
			//{
			//	Name:       "connection, region, service scoped limiter",
			//	FillRate:   10,
			//	BucketSize: 10,
			//	Scopes:     []string{rateLimiterScopeRegion, rateLimiterScopeService, rate_limiter.RateLimiterScopeConnection},
			//},
		},
	}
	//var a map[string]*plugin.Plugin
	//
	//a["foo"] = p
	//return

	return p
}
