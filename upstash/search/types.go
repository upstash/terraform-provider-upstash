package search

type Search struct {
	Id                   string  `json:"id"`
	Name                 string  `json:"name"`
	Endpoint             string  `json:"endpoint"`
	Token                string  `json:"token"`
	ReadOnlyToken        string  `json:"read_only_token"`
	Type                 string  `json:"type"`
	Region               string  `json:"region"`
	CreationTime         int64   `json:"creation_time"`
	MaxVectorCount       int64   `json:"max_vector_count"`
	MaxDailyUpdates      int64   `json:"max_daily_updates"`
	MaxDailyQueries      int64   `json:"max_daily_queries"`
	MaxMonthlyBandwidth  int64   `json:"max_monthly_bandwidth"`
	MaxWritesPerSecond   int64   `json:"max_writes_per_second"`
	MaxQueryPerSecond    int64   `json:"max_query_per_second"`
	MaxReadsPerRequest   int64   `json:"max_reads_per_request"`
	MaxWritesPerRequest  int64   `json:"max_writes_per_request"`
	MaxTotalMetadataSize int64   `json:"max_total_metadata_size"`
	ReservedPrice        float64 `json:"reserved_price"`
	CustomerId           string  `json:"customer_id"`
}

type CreateSearchRequest struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Type   string `json:"type"`
}

type SetPlanRequest struct {
	TargetPlan string `json:"target_plan"`
}

type TransferSearchRequest struct {
	TargetAccount string `json:"target_account"`
}

type RenameSearchRequest struct {
	Name string `json:"name"`
}
