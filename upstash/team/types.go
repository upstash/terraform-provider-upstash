package team

type Team struct {
	TeamId   string `json:"team_id"`
	TeamName string `json:"team_name"`
	// not needed, if changes --> reCreate
	// CopyCC   bool   `json:"copy_cc"`

	// Newly added, what should be done here?
	TeamMembers map[string]string `json:"team_members"`
}

type CreateTeamRequest struct {
	TeamName string `json:"team_name"`
	CopyCC   bool   `json:"copy_cc"`
}

type GetTeamMembers struct {
	TeamId      string `json:"team_id"`
	TeamName    string `json:"team_name"`
	MemberEmail string `json:"member_email"`
	MemberRole  string `json:"member_role"`
}

type Member struct {
	MemberEmail string
	MemberRole  string
}
