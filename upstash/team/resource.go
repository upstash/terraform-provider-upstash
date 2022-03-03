package team

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceTeam() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"team_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique Cluster ID for created cluster",
			},
			"team_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the team",
			},
			"copy_cc": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Whether Credit Card is copied",
			},
			"team_members": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Members of the team. (Owner must be specified, which is the owner of the api key.)",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					noOwner := true
					for _, b := range val.(map[string]interface{}) {
						if b.(string) == "owner" {
							noOwner = false
						}
					}
					if noOwner {
						errs = append(errs, fmt.Errorf("Owner of the api key should be given the role of owner"))
					}
					return
				},
			},
		},
	}

}
