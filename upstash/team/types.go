package team

// lacks some fields --> members?
type Team struct {
	TeamId   string `json:"team_id"`
	TeamName string `json:"team_name"`
	CopyCC   bool   `json:"copy_cc"`
}

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
	CopyCC   bool   `json:"copy_cc"`
}
