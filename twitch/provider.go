package twitch

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"TWITCH_API_KEY",
				}, nil),
				Description: "twitch API Key from https://dev.twitch.tv",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"twitch_user":    dataSourceTwitchUser(),
			"twitch_channel": dataSourceTwitchChannel(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiKey: d.Get("api_key").(string),
	}
	return config.Client()
}
