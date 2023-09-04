package chaosratelimit

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"log"
)

const pluginRateLimiterListRowCount = 500

func rateLimiterTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaosratelimit_rate_limiter",
		Description: "Table which uses the plugin-scoped rate limiter_",
		List: &plugin.ListConfig{
			Hydrate: pluginRateLimiterList,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:      "region",
					Operators: []string{"="},
					Require:   plugin.Optional,
				},
				{
					Name:      "qual",
					Operators: []string{"="},
					Require:   plugin.Optional,
				},
			},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: rate4Hydrate,
			},
			{
				Func: rate3Hydrate,
			},
			{
				Func: rate2Hydrate,
			},
			{
				Func: GetTopicAttributes,
				Tags: map[string]string{"service": "sns", "action": "GetTopicAttributes"},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT},
			{Name: "region", Type: proto.ColumnType_STRING},
			{Name: "qual", Type: proto.ColumnType_STRING},

			{Name: "rate_4", Type: proto.ColumnType_INT, Hydrate: rate4Hydrate},
			{Name: "rate_3", Type: proto.ColumnType_INT, Hydrate: rate3Hydrate},
			{Name: "rate_2", Type: proto.ColumnType_INT, Hydrate: rate2Hydrate},
			{Name: "rate_1", Type: proto.ColumnType_INT, Hydrate: GetTopicAttributes},
		},
		GetMatrixItemFunc: getRegions,
	}
}

func getRegions(context.Context, *plugin.QueryData) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"region": "us-east-1",
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
func GetTopicAttributes(context.Context, *plugin.QueryData, *plugin.HydrateData) (interface{}, error) {
	return map[string]string{"rate_1": "1"}, nil

}

func pluginRateLimiterList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < pluginRateLimiterListRowCount; i++ {
		item := populateItem(i, d.Table)

		if region, ok := d.EqualsQuals["region"]; ok {
			item["region"] = region.GetStringValue()
		}
		if qual, ok := d.EqualsQuals["qual"]; ok && item["qual"] != qual.GetStringValue() {
			continue
		}

		d.StreamListItem(ctx, item)
		if d.RowsRemaining(ctx) == 0 {
			log.Printf("[WARN] HIT LIMIT, EXITING")
			break
		}
	}

	return nil, nil
}
