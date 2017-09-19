package twitch

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceTwitchUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTwitchUserRead,

		Schema: map[string]*schema.Schema{
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTwitchUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*twitchClient).conn
	me, err := client.GetUser(nil)
	if err != nil {
		log.Fatalf("Error finding me: %s", err)
	}

	d.Set("user_id", me.Id)
	d.Set("display_name", me.DisplayName)
	d.SetId(strconv.Itoa(me.Id))
	return nil
}
