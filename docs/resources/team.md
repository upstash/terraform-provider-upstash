---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "upstash_team Resource - terraform-provider-upstash"
subcategory: ""
description: |-
  
---

# upstash_team (Resource)



## Example Usage

```terraform
resource "upstash_team" "exampleTeam" {
    team_name = "TerraformTeam"
    copy_cc = false

    team_members = {
        # Owner is the owner of the api_key.
        "X@Y.Z": "owner",
        "A@B.C": "dev",
        "E@E.F": "finance",
    }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `copy_cc` (Boolean) Whether Credit Card is copied
- `team_members` (Map of String) Members of the team. (Owner must be specified, which is the owner of the api key.)
- `team_name` (String) Name of the team

### Read-Only

- `id` (String) The ID of this resource.
- `team_id` (String) Unique Cluster ID for created cluster
