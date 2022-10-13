package schedule

import "github.com/upstash/terraform-provider-upstash/upstash/qstash/topic"

type CreateQstashScheduleRequest struct {
	Destination      string           `json:"destination"`
	HeaderParameters HeaderParameters `json:"headerParameters"`
}

type QstashSchedule struct {
	Content     Content     `json:"content"`
	CreatedAt   int64       `json:"createdAt,omitempty"`
	Cron        string      `json:"cron"`
	Destination Destination `json:"destination"`
	ScheduleId  string      `json:"scheduleId"`
	Settings    Settings    `json:"settings"`
}

type Settings struct {
	Deadline  int64 `json:"deadline"`
	NotBefore int64 `json:"notBefore"`
	Retries   int32 `json:"retries"`
}

type Destination struct {
	Topic topic.QstashTopic `json:"topic,omitempty"`
	Type  string            `json:"type"`
	Url   string            `json:"url,omitempty"`
}

type Content struct {
	Body   []int             `json:"body"`
	Header map[string]string `json:"header"`
}

type HeaderParameters struct {
	ContentType                      string
	UpstashDeduplicationId           string
	UpstashContentBasedDeduplication string
	UpstashNotBefore                 string
	UpstashDelay                     string
	UpstashRetries                   string
	UpstashCron                      string
	// UpstashContentBasedDeduplication bool
	// UpstashNotBefore                 int64
	// UpstashRetries                   int64
}
