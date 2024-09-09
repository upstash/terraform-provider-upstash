package database

type Database struct {
	DatabaseId             string   `json:"database_id"`
	DatabaseName           string   `json:"database_name"`
	Region                 string   `json:"region"`
	Port                   int      `json:"port"`
	CreationTime           int64    `json:"creation_time"`
	Password               string   `json:"password,omitempty"`
	Endpoint               string   `json:"endpoint"`
	Tls                    bool     `json:"tls"`
	Eviction               bool     `json:"eviction"`
	AutoUpgrade            bool     `json:"auto_upgrade"`
	Consistent             bool     `json:"consistent"`
	MultiZone              bool     `json:"multizone"`
	RestToken              string   `json:"rest_token,omitempty"`
	ReadOnlyRestToken      string   `json:"read_only_rest_token,omitempty"`
	DatabaseType           string   `json:"database_type"`
	State                  string   `json:"state"`
	UserEmail              string   `json:"user_email"`
	DBMaxClients           int      `json:"db_max_clients"`
	DBMaxRequestSize       int64    `json:"db_max_request_size"`
	DBDiskThreshold        int64    `json:"db_disk_threshold"`
	DBMaxEntrySize         int64    `json:"db_max_entry_size"`
	DBMemoryThreshold      int64    `json:"db_memory_threshold"`
	DBDailyBandwidthLimit  int64    `json:"db_daily_bandwidth_limit"`
	DBMaxCommandsPerSecond int64    `json:"db_max_commands_per_second"`
	PrimaryRegion          string   `json:"primary_region"`
	ReadRegions            []string `json:"read_regions"`
}

type CreateDatabaseRequest struct {
	Region        string   `json:"region"`
	DatabaseName  string   `json:"database_name"`
	Eviction      bool     `json:"eviction"`
	AutoUpgrade   bool     `json:"auto_upgrade"`
	PrimaryRegion string   `json:"primary_region,omitempty"`
	ReadRegions   []string `json:"read_regions,omitempty"`
}

type ConfigureEvictionRequest struct {
	DatabaseId string `json:"database_id"`
	Eviction   bool   `json:"eviction"`
}

type ConfigureAutoUpgradeRequest struct {
	DatabaseId  string `json:"database_id"`
	AutoUpgrade bool   `json:"auto_upgrade"`
}

type UpdateReadRegionsRequest struct {
	ReadRegions []string `json:"read_regions"`
}
