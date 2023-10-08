package schedule

type QstashSchedule struct {
	CreatedAt   int64               `json:"createdAt"`
	ScheduleId  string              `json:"scheduleId"`
	Cron        string              `json:"cron"`
	Destination string              `json:"destination"`
	Method      string              `json:"method"`
	Header      map[string][]string `json:"header,omitempty"`
	Body        string              `json:"body,omitempty"`
	Retries     int                 `json:"retries"`
	Delay       int                 `json:"delay,omitempty"`
	Callback    string              `json:"callback,omitempty"`
}

type CreateQstashScheduleResponse struct {
	ScheduleId string `json:"scheduleId"`
}

type CreateQstashScheduleRequest struct {
	Destination    string `json:"destination"`
	Body           string
	ForwardHeaders map[string]interface{}
	Headers        QstashScheduleHeaders
}

type QstashScheduleHeaders struct {
	ContentType string
	Method      string
	Delay       string
	Retries     int
	Callback    string
	Cron        string
}
