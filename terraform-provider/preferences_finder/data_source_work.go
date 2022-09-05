package preferences_finder

import (
	"context"
	"log"

	"os"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	firebase "firebase.google.com/go"

	pflib "github.com/sweeneyb/preferences-finder/go-admin/pflib"
)

func dataSourceWorks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorksRead,
		Schema: map[string]*schema.Schema{
			"works": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"citation": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"artist": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWorksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	conf := &firebase.Config{ProjectID: os.Getenv("project_id")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fsclient := pflib.Client{client}
	collection := fsclient.GetWorks("first", ctx)
	var diags diag.Diagnostics
	d.Set("works", flattenWorksData(collection))

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenWorksData(works []pflib.Work) []interface{} {
	if works != nil {
		ois := make([]interface{}, len(works), len(works))

		for i, work := range works {
			oi := make(map[string]interface{})

			oi["id"] = work.ID
			oi["name"] = work.Name
			oi["citation"] = work.Citation
			oi["image_url"] = work.ImageURL
			oi["artist"] = work.Artist

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}
