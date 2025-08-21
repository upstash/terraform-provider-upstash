package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upstash/terraform-provider-upstash/v2/upstash/client"
	"github.com/upstash/terraform-provider-upstash/v2/upstash/utils"
)

func resourceDatabaseUpdate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	databaseId := data.Get("database_id").(string)

	if data.HasChange("read_regions") {
		var readRegions []string
		primaryRegion := data.Get("primary_region").(string)

		for _, v := range (data.Get("read_regions").(*schema.Set)).List() {
			if v.(string) == primaryRegion {
				return diag.Errorf(fmt.Sprintf("Primary region '%s' can not be in the list of read regions.", primaryRegion))
			}
			readRegions = append([]string{v.(string)}, readRegions...)
		}
		if err := UpdateReadRegions(c, databaseId, UpdateReadRegionsRequest{readRegions}); err != nil {
			return diag.FromErr((err))
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

	if data.HasChange("prod_pack") {
		if err := ConfigureProdPack(c, databaseId, data.Get("prod_pack").(bool)); err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("budget") {
		err := UpdateDBBudget(c, databaseId, UpdateDBBudgetRequest{
			Budget: data.Get("budget").(int),
		})

		if err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("ip_allowlist") {
		var ipAllowList []string
		for _, v := range (data.Get("ip_allowlist").(*schema.Set)).List() {
			if v != nil {
				ipAllowList = append(ipAllowList, v.(string))
			}
		}

		err := UpdateDBIpAllowlist(c, databaseId, UpdateDBIpAllowlistRequest{
			AllowedIps: ipAllowList,
		})

		if err != nil {
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

func resourceDatabaseRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)
	databaseId := data.Get("database_id").(string)
	if databaseId == "" {
		databaseId = data.Id()
	}
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
		"prod_pack":                  database.ProdPack,
		"budget":                     database.Budget,
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
		"primary_region":             database.PrimaryRegion,
	}
	if len(database.IpAllowList) > 0 {
		mapping["ip_allowlist"] = database.IpAllowList
	}

	if len(database.ReadRegions) > 0 {
		mapping["read_regions"] = database.ReadRegions
	}

	return utils.SetAndCheckErrors(data, mapping)
}

func resourceDatabaseCreate(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.UpstashClient)

	var readRegions []string
	for _, v := range (data.Get("read_regions").(*schema.Set)).List() {
		if v != nil {
			readRegions = append(readRegions, v.(string))
		}
	}

	database, err := CreateDatabase(c, CreateDatabaseRequest{
		Region:        data.Get("region").(string),
		DatabaseName:  data.Get("database_name").(string),
		Eviction:      data.Get("eviction").(bool),
		AutoUpgrade:   data.Get("auto_scale").(bool),
		ProdPack:      data.Get("prod_pack").(bool),
		Budget:        data.Get("budget").(int),
		PrimaryRegion: data.Get("primary_region").(string),
		Tls:           data.Get("tls").(bool),
		ReadRegions:   readRegions,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("upstash-database-" + database.DatabaseId)
	data.Set("database_id", database.DatabaseId)

	var ipAllowList []string
	for _, v := range (data.Get("ip_allowlist").(*schema.Set)).List() {
		if v != nil {
			ipAllowList = append(ipAllowList, v.(string))
		}
	}

	if len(ipAllowList) > 0 {
		err = UpdateDBIpAllowlist(c, database.DatabaseId, UpdateDBIpAllowlistRequest{
			AllowedIps: ipAllowList,
		})
		if err != nil {
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
