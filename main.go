package main

import (
	"github.com/turbot/steampipe-plugin-chaos/chaosratelimit"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaosratelimit.Plugin})
}
