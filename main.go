package main

import (
	"github.com/catsby/terraform-provider-twitch/twitch"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: twitch.Provider})
}
