package integrationtesting

import (
	"os"
	"strconv"
	"strings"
)

var email, apikey string

type EnvVars struct {
	Email  string
	Apikey string

	RedisDatabaseName      string
	RedisDatabaseRegion    string
	RedisDatabaseTls       bool
	RedisDatabaseMultiZone bool

	VectorIndexName               string
	VectorIndexType               string
	VectorIndexRegion             string
	VectorIndexDimensionCount     int
	VectorIndexSimilarityFunction string

	TeamName    string
	CopyCC      bool
	TeamMembers map[string]string
}

func GetEnvVars() EnvVars {
	vectorIndexDimensionCount, _ := strconv.Atoi(os.Getenv("UPSTASH_VECTOR_INDEX_DIMENSION_COUNT"))

	teamMembers := make(map[string]string)

	teamOwner := strings.Fields(os.Getenv("UPSTASH_TEAM_OWNER"))
	if len(teamOwner) != 0 {
		teamMembers[teamOwner[0]] = "owner"
	}

	teamDevs := strings.Fields(os.Getenv("UPSTASH_TEAM_DEVS"))
	teamFinances := strings.Fields(os.Getenv("UPSTASH_TEAM_FINANCES"))

	for _, val := range teamDevs {
		teamMembers[val] = "dev"
	}

	for _, val := range teamFinances {
		teamMembers[val] = "finance"
	}

	return EnvVars{
		Email:  os.Getenv("UPSTASH_EMAIL"),
		Apikey: os.Getenv("UPSTASH_API_KEY"),

		RedisDatabaseName:      os.Getenv("UPSTASH_REDIS_DATABASE_NAME"),
		RedisDatabaseRegion:    os.Getenv("UPSTASH_REDIS_DATABASE_REGION"),
		RedisDatabaseTls:       os.Getenv("UPSTASH_REDIS_DATABASE_TLS") == "true",
		RedisDatabaseMultiZone: os.Getenv("UPSTASH_REDIS_DATABASE_MULTIZONE") == "true",

		VectorIndexName:               os.Getenv("UPSTASH_VECTOR_INDEX_NAME"),
		VectorIndexType:               os.Getenv("UPSTASH_VECTOR_INDEX_TYPE"),
		VectorIndexRegion:             os.Getenv("UPSTASH_VECTOR_INDEX_REGION"),
		VectorIndexDimensionCount:     vectorIndexDimensionCount,
		VectorIndexSimilarityFunction: os.Getenv("UPSTASH_VECTOR_INDEX_SIMILARITY_FUNCTION"),

		TeamName:    os.Getenv("UPSTASH_TEAM_NAME"),
		CopyCC:      os.Getenv("UPSTASH_TEAM_COPY_CC") == "true",
		TeamMembers: teamMembers,
	}
}
