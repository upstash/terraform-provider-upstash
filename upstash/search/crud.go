package search

import (
	"context"

	"github.com/upstash/terraform-provider-upstash/v2/upstash/utils"

	"github.com/upstash/terraform-provider-upstash/v2/upstash/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSearchCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)

	search, err := CreateSearch(c, CreateSearchRequest{
		Name:   data.Get("name").(string),
		Region: data.Get("region").(string),
		Type:   data.Get("type").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(search.Id)

	err = data.Set("id", search.Id)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSearchRead(ctx, data, m)
}

func resourceSearchRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	searchId := data.Get("id").(string)

	if searchId == "" {
		searchId = data.Id()
	}

	search, err := GetSearch(c, searchId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(search.Id)

	searchType := search.Type
	if searchType == "paid" {
		searchType = "payg"
	}

	mapping := map[string]interface{}{
		"customer_id":             search.CustomerId,
		"id":                      search.Id,
		"name":                    search.Name,
		"endpoint":                search.Endpoint,
		"token":                   search.Token,
		"read_only_token":         search.ReadOnlyToken,
		"type":                    searchType,
		"region":                  search.Region,
		"max_vector_count":        search.MaxVectorCount,
		"max_daily_updates":       search.MaxDailyUpdates,
		"max_daily_queries":       search.MaxDailyQueries,
		"max_monthly_bandwidth":   search.MaxMonthlyBandwidth,
		"max_writes_per_second":   search.MaxWritesPerSecond,
		"max_query_per_second":    search.MaxQueryPerSecond,
		"max_reads_per_request":   search.MaxReadsPerRequest,
		"max_writes_per_request":  search.MaxWritesPerRequest,
		"max_total_metadata_size": search.MaxTotalMetadataSize,
		"reserved_price":          search.ReservedPrice,
		"creation_time":           search.CreationTime,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceSearchUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	searchId := data.Get("id").(string)

	if data.HasChange("type") {
		err := SetSearchPlan(c, searchId, SetPlanRequest{TargetPlan: data.Get("type").(string)})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("name") {
		err := RenameSearch(c, searchId, RenameSearchRequest{Name: data.Get("name").(string)})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceSearchRead(ctx, data, m)
}

func resourceSearchDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	searchId := data.Get("id").(string)

	err := DeleteSearch(c, searchId)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
