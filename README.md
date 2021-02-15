# PRCheckCtl
Simple tool to check for new PRs by username or organization on Github

> Just for practice Golang!

# Usage
1. Build the binary (inside the code directory):   
`go build`

2. Export your github token: 
`export GH_TOKEN=<your_github_token>`

3. Run the binary:   
`./prcheckctl --username <github_username or org> --pool <time to check for new PRs>`

## Examples

1. Check for new PRs in my repositories using 15 seconds between every check:   
`prcheckctl --username edsoncelio--pool 15`

To get help, use:
`prcheckctl --help`

# TODO
- [ ] Tests
- [ ] Docs
- [ ] Integrations
- [ ] Refactor some erros and treat exceptions!
