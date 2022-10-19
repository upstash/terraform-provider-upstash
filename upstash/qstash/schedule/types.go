package schedule

import "github.com/upstash/terraform-provider-upstash/upstash/qstash/topic"

type CreateQstashScheduleRequest struct {
	Destination    string `json:"destination"`
	Body           string
	ForwardHeaders map[string]interface{}
	Headers        QstashScheduleHeaders
}

type QstashSchedule struct {
	Content     Content     `json:"content"`
	CreatedAt   int64       `json:"createdAt"`
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
	Body   string            `json:"body"`
	Header map[string]string `json:"header"`
}

type QstashScheduleHeaders struct {
	ContentType               string
	DeduplicationId           string
	ContentBasedDeduplication bool
	NotBefore                 int
	Delay                     string
	Retries                   int
	Cron                      string
	ForwardHeaders            map[string]string
}
