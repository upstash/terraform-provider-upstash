resource "upstash_team" "exampleTeam" {
    team_name = "TerraformTeam"
    copy_cc = false
    team_members = {
        # Owner is the owner of the api_key.
        "X@Y.Z": "owner",
        "A@B.C": "dev",
        "D@E.F": "finance",
    }
}