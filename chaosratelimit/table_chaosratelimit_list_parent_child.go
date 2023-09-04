package chaosratelimit

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type parentData struct {
	Id int
}

type childData struct {
	Id          int
	ChildColumn string
}

func parentChildTable() *plugin.Table {
	return &plugin.Table{
		Name:        "chaosratelimit_list_parent_child",
		Description: "Chaos table to test the List calls having parent-child dependencies with all the possible scenarios like errors, panics and delays at both parent and child levels",
		List: &plugin.ListConfig{
			Hydrate:       listChildHydrateTable,
			ParentHydrate: listParentHydrateTable,
			Tags:          map[string]string{"limiter_rate": "4"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("Id"),
			Hydrate:    getTable,
			Tags:       map[string]string{"limiter_rate": "1"},
		},
		Columns: []*plugin.Column{
			{Name: "Id", Type: proto.ColumnType_INT, Description: "Column for the ID"},
			{Name: "ChildColumn", Type: proto.ColumnType_STRING, Description: "Column to test the the parent list function with fatal error after streaming some rows"},
		},
	}
}

func listParentHydrateTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	for i := 0; i < 10; i++ {
		item := &parentData{i}
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}
func listChildHydrateTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	parent := h.Item.(*parentData)

	for i := 0; i < 10; i++ {
		item := &childData{Id: parent.Id*10 + i, ChildColumn: fmt.Sprintf("child_column-%d", i)}
		d.StreamListItem(ctx, item)
	}

	return nil, nil

}

func getTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetInt64Value()

	item := map[string]interface{}{"id": id, "child_column": fmt.Sprintf("child_column-get-%v", id)}
	d.StreamListItem(ctx, item)
	return nil, nil
}
