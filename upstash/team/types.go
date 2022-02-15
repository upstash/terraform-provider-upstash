package team

// lacks some fields --> members?
type Team struct {
	TeamId   string `json:"team_id"`
	TeamName string `json:"team_name"`
	CopyCC   bool   `json:"copy_cc"`

	// Newly added, what should be done here?
	Members []string `json:"members"`
}

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
	CopyCC   bool   `json:"copy_cc"`
}
