output "team_id" {
  value = data.upstash_team_data.teamData.team_id
}

output "team_name" {
  value = data.upstash_team_data.teamData.team_name
}

output "team_members" {
  value = data.upstash_team_data.teamData.team_members
}

output "copy_cc" {
  value = resource.upstash_team.exampleTeam.copy_cc
}

output "copy_cc_from_data" {
  value = data.upstash_team_data.teamData.copy_cc
}
