package chaosratelimit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"log"
)

const pluginRateLimiterListRowCount = 10000

func rateLimiterTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaosratelimit_rate_limiter",
		Description: "Table which uses the plugin-scoped rate limiter",
		List: &plugin.ListConfig{
			Hydrate: pluginRateLimiterList,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:        rate4Hydrate,
				ScopeValues: map[string]string{"limiter_rate": "4"},
			},
			{
				Func:        rate3Hydrate,
				ScopeValues: map[string]string{"limiter_rate": "3"},
			},
			{
				Func:        rate2Hydrate,
				ScopeValues: map[string]string{"limiter_rate": "2"},
			},
			{
				Func:        rate1Hydrate,
				ScopeValues: map[string]string{"limiter_rate": "1"},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "rate_4", Type: proto.ColumnType_INT, Hydrate: rate4Hydrate},
			{Name: "rate_3", Type: proto.ColumnType_INT, Hydrate: rate3Hydrate},
			{Name: "rate_2", Type: proto.ColumnType_INT, Hydrate: rate2Hydrate},
			{Name: "rate_1", Type: proto.ColumnType_INT, Hydrate: rate1Hydrate},
		},
	}

}

func rate4Hydrate(context.Context, *plugin.QueryData, *plugin.HydrateData) (interface{}, error) {
	return map[string]string{"rate_4": "4"}, nil
}
func rate3Hydrate(context.Context, *plugin.QueryData, *plugin.HydrateData) (interface{}, error) {
	return map[string]string{"rate_3": "3"}, nil

}
func rate2Hydrate(context.Context, *plugin.QueryData, *plugin.HydrateData) (interface{}, error) {
	return map[string]string{"rate_2": "2"}, nil

}
func rate1Hydrate(context.Context, *plugin.QueryData, *plugin.HydrateData) (interface{}, error) {
	return map[string]string{"rate_1": "1"}, nil

}

func pluginRateLimiterList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < pluginRateLimiterListRowCount; i++ {
		d.StreamListItem(ctx, populateItem(i, d.Table))
		if d.RowsRemaining(ctx) == 0 {
			log.Printf("[WARN] HIT LIMIT, EXITING")
			break
		}
	}

	return nil, nil
}
