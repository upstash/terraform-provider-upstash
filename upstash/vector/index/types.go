package index

type Index struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	SimilarityFunction   string `json:"similarity_function"`
	DimensionCount       int64  `json:"dimension_count"`
	Endpoint             string `json:"endpoint"`
	Token                string `json:"token"`
	ReadOnlyToken        string `json:"read_only_token"`
	Type                 string `json:"type"`
	Region               string `json:"region"`
	CreationTime         int64  `json:"creation_time"`
	MaxVectorCount       int64  `json:"max_vector_count"`
	MaxDailyUpdates      int64  `json:"max_daily_updates"`
	MaxDailyQueries      int64  `json:"max_daily_queries"`
	MaxMonthlyBandwidth  int64  `json:"max_monthly_bandwidth"`
	MaxWritesPerSecond   int64  `json:"max_writes_per_second"`
	MaxQueryPerSecond    int64  `json:"max_query_per_second"`
	MaxReadsPerRequest   int64  `json:"max_reads_per_request"`
	MaxWritesPerRequest  int64  `json:"max_writes_per_request"`
	MaxTotalMetadataSize int64  `json:"max_total_metadata_size"`
	ReservedPrice        int64  `json:"reserved_price"`
	CustomerId           string `json:"customer_id"`
}

type CreateIndexRequest struct {
	Name               string `json:"name"`
	SimilarityFunction string `json:"similarity_function"`
	DimensionCount     int64  `json:"dimension_count"`
	Region             string `json:"region"`
	Type               string `json:"type"`
}

type SetPlanRequest struct {
	TargetPlan string `json:"target_plan"`
}

type TransferIndexRequest struct {
	TargetAccount string `json:"target_account"`
}
