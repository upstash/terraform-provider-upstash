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

	if data.HasChange("tls") {
		if err := EnableTLS(c, databaseId); err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("eviction") {
		if err := ConfigureEviction(c, databaseId, data.Get("eviction").(bool)); err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("auto_scale") {
		if err := ConfigureAutoUpgrade(c, databaseId, data.Get("auto_scale").(bool)); err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("consistent") {
		if data.Get("consistent").(bool) {
			return diag.Errorf("Cannot enable strong consistency on the DB. All the newly created DBs will be eventually consistent. Set consistent=false.")
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
		"database_id":                database.DatabaseId,
		"database_name":              database.DatabaseName,
		"region":                     database.Region,
		"endpoint":                   database.Endpoint,
		"password":                   database.Password,
		"consistent":                 database.Consistent,
		"multizone":                  database.MultiZone,
		"tls":                        database.Tls,
		"eviction":                   database.Eviction,
		"auto_scale":                 database.AutoUpgrade,
		"port":                       database.Port,
		"rest_token":                 database.RestToken,
		"read_only_rest_token":       database.ReadOnlyRestToken,
		"database_type":              database.DatabaseType,
		"state":                      database.State,
		"user_email":                 database.UserEmail,
		"db_max_clients":             database.DBMaxClients,
		"db_max_request_size":        database.DBMaxRequestSize,
		"db_disk_threshold":          database.DBDiskThreshold,
		"db_max_entry_size":          database.DBMaxEntrySize,
		"db_memory_threshold":        database.DBMemoryThreshold,
		"db_daily_bandwidth_limit":   database.DBDailyBandwidthLimit,
		"db_max_commands_per_second": database.DBMaxCommandsPerSecond,
		"creation_time":              database.CreationTime,
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceDatabaseCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	if data.Get("consistent").(bool) {
		return diag.Errorf("Newly created DBs are eventually consistent. Set consistent=false in the resource.")
	}
	c := m.(*client.UpstashClient)
	database, err := CreateDatabase(c, CreateDatabaseRequest{
		Region:       data.Get("region").(string),
		DatabaseName: data.Get("database_name").(string),
		Tls:          data.Get("tls").(bool),
		Eviction:     data.Get("eviction").(bool),
		AutoUpgrade:  data.Get("auto_scale").(bool),
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
