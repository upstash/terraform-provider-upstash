package index

import (
	"context"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"

	//"fmt"
	"github.com/upstash/terraform-provider-upstash/upstash/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	_ "github.com/upstash/terraform-provider-upstash/upstash/client"
	_ "github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceIndexCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)

	index, err := CreateIndex(c, CreateIndexRequest{
		Name:               data.Get("name").(string),
		SimilarityFunction: data.Get("similarity_function").(string),
		DimensionCount:     data.Get("dimension_count").(int64),
		Region:             data.Get("region").(string),
		Type:               data.Get("type").(string),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-index-" + index.Id)

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

	mapping := map[string]interface{}{
		"customer_id":             index.CustomerId,
		"id":                      index.Id,
		"name":                    index.Name,
		"similarity_function":     index.SimilarityFunction,
		"dimension_count":         index.DimensionCount,
		"endpoint":                index.Endpoint,
		"token":                   index.Token,
		"read_only_token":         index.ReadOnlyToken,
		"type":                    index.Type,
		"max_vector_count":        index.MaxVectorCount,
		"max_daily_updates":       index.MaxDailyUpdates,
		"max_daily_queries":       index.MaxDailyQueries,
		"max_monthly_bandwidth":   index.MaxMonthlyBandwidth,
		"max_writes_per_second":   index.MaxWritesPerSecond,
		"max_query_per_second":    index.MaxQueryPerSecond,
		"max_reads_per_seconds":   index.MaxReadsPerRequest,
		"max_total_metadata_size": index.MaxTotalMetadataSize,
		"reserved_price":          index.ReservedPrice,
		"creation_time":           index.CreationTime,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceIndexUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	indexId := data.Get("indexId").(string)

	if data.HasChange("plan") {
		if err := SetIndexPlan(c, indexId, data.Get("plan").(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("name") {
		if err := RenameIndex(c, indexId, data.Get("name").(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceIndexRead(ctx, data, m)
}

func resourceIndexDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	indexId := data.Get("indexId").(string)

	err := DeleteIndex(c, indexId)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
