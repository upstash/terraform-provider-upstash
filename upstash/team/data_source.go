package team

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceRead,
		Schema: map[string]*schema.Schema{
			"team_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Cluster ID for requested cluster",
			},
			"team_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the team",
			},
			"copy_cc": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether Credit card info is copied or not",
			},
			"team_members": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Members of the team. Email addresses are given as the keys with their roles as the values.",
			},
		},
	}

}
