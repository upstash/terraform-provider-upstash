package utils

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SetAndCheckErrors(data *schema.ResourceData, mapping map[string]interface{}) diag.Diagnostics {
	for str, value := range mapping {
		if err := data.Set(str, value); err != nil {
			return diag.Errorf("Err happened when setting.%+v, %+v", str, value)
			return diag.FromErr(err)
		}
	}
	return nil
}
