package twitch

import (
	"log"
	"strconv"

	"github.com/catsby/go-twitch/twitch"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceTwitchChannel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTwitchChannelRead,

		Schema: map[string]*schema.Schema{
			"channel_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTwitchChannelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*twitchClient).conn
	chInput := twitch.GetChannelInput{}
	if v, ok := d.GetOk("channel_id"); ok {
		id, err := strconv.Atoi(v.(string))
		if err != nil {
			log.Printf("error parsing channel_id")
		} else {
			chInput.Id = id
		}
	}
	output, err := client.GetChannel(&chInput)
	if err != nil {
		log.Fatalf("Error finding channel: %s", err)
	}

	log.Printf("what is channel: %s", spew.Sdump(output))
	log.Printf("got here")
	d.Set("channel_id", output.Channel.Id)
	d.Set("display_name", output.Channel.DisplayName)
	d.Set("url", output.Channel.HTMLURL)
	d.SetId(strconv.Itoa(output.Channel.Id))
	return nil
}
