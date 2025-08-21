package index

import (
	"context"

	"github.com/upstash/terraform-provider-upstash/v2/upstash/utils"

	"github.com/upstash/terraform-provider-upstash/v2/upstash/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIndexCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)

	index, err := CreateIndex(c, CreateIndexRequest{
		Name:               data.Get("name").(string),
		SimilarityFunction: data.Get("similarity_function").(string),
		DimensionCount:     data.Get("dimension_count").(int),
		Region:             data.Get("region").(string),
		Type:               data.Get("type").(string),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(index.Id)

	err = data.Set("id", index.Id)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIndexRead(ctx, data, m)
}

func resourceIndexRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	indexId := data.Get("id").(string)

	if indexId == "" {
		indexId = data.Id()
	}

	index, err := GetIndex(c, indexId)

	if err != nil {
		return diag.FromErr(err)
	}
	indexType := index.Type

	if indexType == "paid" {
		indexType = "payg"
	}

	mapping := map[string]interface{}{
		"customer_id":             index.CustomerId,
		"id":                      index.Id,
		"name":                    index.Name,
		"similarity_function":     index.SimilarityFunction,
		"dimension_count":         index.DimensionCount,
		"endpoint":                index.Endpoint,
		"token":                   index.Token,
		"read_only_token":         index.ReadOnlyToken,
		"type":                    indexType,
		"region":                  index.Region,
		"max_vector_count":        index.MaxVectorCount,
		"max_daily_updates":       index.MaxDailyUpdates,
		"max_daily_queries":       index.MaxDailyQueries,
		"max_monthly_bandwidth":   index.MaxMonthlyBandwidth,
		"max_writes_per_second":   index.MaxWritesPerSecond,
		"max_query_per_second":    index.MaxQueryPerSecond,
		"max_reads_per_request":   index.MaxReadsPerRequest,
		"max_writes_per_request":  index.MaxWritesPerRequest,
		"max_total_metadata_size": index.MaxTotalMetadataSize,
		"reserved_price":          index.ReservedPrice,
		"creation_time":           index.CreationTime,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceIndexUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	indexId := data.Get("id").(string)

	if data.HasChange("type") {
		err := SetIndexPlan(c, indexId, SetPlanRequest{TargetPlan: data.Get("type").(string)})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("name") {
		err := RenameIndex(c, indexId, RenameIndexRequest{Name: data.Get("name").(string)})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceIndexRead(ctx, data, m)
}

func resourceIndexDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	indexId := data.Get("id").(string)

	err := DeleteIndex(c, indexId)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
