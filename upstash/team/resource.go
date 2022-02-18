package team

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
				Description: "Name of the cluster",
			},
			"copy_cc": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Whether Credit Card is copied",
			},
			"team_members": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Members of the team",
			},
		},
	}

}
