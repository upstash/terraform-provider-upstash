package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/upstash/client"
	"github.com/upstash/terraform-provider-upstash/upstash/utils"
)

func resourceDatabaseUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	databaseId := data.Get("database_id").(string)
	if data.HasChange("multizone") {
		if err := EnableMultiZone(c, databaseId, data.Get("multizone").(bool)); err != nil {
			return diag.FromErr(err)
		}
	}
	// I think this creates a problem. Doesnt give tls as a parameter for the client to use.
	if data.HasChange("tls") {
		if err := EnableTLS(c, databaseId); err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceDatabaseRead(ctx, data, m)
}

func resourceDatabaseDelete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	databaseId := data.Get("database_id").(string)
	err := DeleteDatabase(c, databaseId)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceDatabaseRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	databaseId := data.Get("database_id").(string)
	database, err := GetDatabase(c, databaseId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("upstash-database-" + database.DatabaseId)

	mapping := map[string]interface{}{
		"database_id":   database.DatabaseId,
		"database_name": database.DatabaseName,
		"region":        database.Region,
		"endpoint":      database.Endpoint,
		"password":      database.Password,
		"consistent":    database.Consistent,
		"multizone":     database.MultiZone,
		"tls":           database.Tls,
	}

	return utils.SetAndCheckErrors(data, mapping)

}

func resourceDatabaseCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	database, err := CreateDatabase(c, CreateDatabaseRequest{
		Region:       data.Get("region").(string),
		DatabaseName: data.Get("database_name").(string),
		Tls:          data.Get("tls").(bool),
		Consistent:   data.Get("consistent").(bool),
		MultiZone:    data.Get("multizone").(bool),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-database-" + database.DatabaseId)
	data.Set("database_id", database.DatabaseId)
	return resourceDatabaseRead(ctx, data, m)
}
