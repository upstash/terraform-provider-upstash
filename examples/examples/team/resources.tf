resource "upstash_team" "exampleTeam" {
    team_name = var.team_name
    copy_cc = var.copy_cc
    team_members = var.team_members
}

# resource "upstash_team" "importTeam" {}
