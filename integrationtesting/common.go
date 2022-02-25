package integrationtesting

import (
	"os"
	"strconv"
	"strings"
)

type EnvVars struct {
	Email  string
	Apikey string

	RedisDatabaseName      string
	RedisDatabaseRegion    string
	RedisDatabaseTls       bool
	RedisDatabaseMultiZone bool

	KafkaClusterName      string
	KafkaClusterRegion    string
	KafkaClusterMultiZone bool

	KafkaTopicName           string
	KafkaTopicPartitions     int
	KafkaTopicRetentionTime  int
	KafkaTopicRetentionSize  int
	KafkaTopicMaxMessageSize int
	KafkaTopicCleanupPolicy  string

	TeamName    string
	CopyCC      bool
	TeamMembers map[string]string
}

func GetEnvVars() EnvVars {
	a, _ := strconv.Atoi(os.Getenv("UPSTASH_KAFKA_TOPIC_PARTITIONS"))
	b, _ := strconv.Atoi(os.Getenv("UPSTASH_KAFKA_TOPIC_RETENTION_TIME"))
	c, _ := strconv.Atoi(os.Getenv("UPSTASH_KAFKA_TOPIC_RETENTION_SIZE"))
	d, _ := strconv.Atoi(os.Getenv("UPSTASH_KAFKA_TOPIC_MAX_MESSAGE_SIZE"))

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

		KafkaClusterName:      os.Getenv("UPSTASH_KAFKA_CLUSTER_NAME"),
		KafkaClusterRegion:    os.Getenv("UPSTASH_KAFKA_CLUSTER_REGION"),
		KafkaClusterMultiZone: os.Getenv("UPSTASH_KAFKA_CLUSTER_MULTIZONE") == "true",

		KafkaTopicName:           os.Getenv("UPSTASH_KAFKA_TOPIC_NAME"),
		KafkaTopicPartitions:     a,
		KafkaTopicRetentionTime:  b,
		KafkaTopicRetentionSize:  c,
		KafkaTopicMaxMessageSize: d,
		KafkaTopicCleanupPolicy:  os.Getenv("UPSTASH_KAFKA_TOPIC_CLEANUP_POLICY"),

		TeamName:    os.Getenv("UPSTASH_TEAM_NAME"),
		CopyCC:      os.Getenv("UPSTASH_TEAM_COPY_CC") == "true",
		TeamMembers: teamMembers,
	}
}
