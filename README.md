# Fat Hippo: a revival of old ambitions

## To Dos (in no particular order):
- [ ] add pre commit hook -> for linting and testing
- [ ] finish hangman full gam loop
- [ ] accept player guesses
- [ ] validate player inputs
- [ ] add tests for everything
- [ ] connect shell package to hangman
- [ ] create a build target
- [ ] create a CI workflow
- [ ] create a remote repo and open source it
- [ ] add juice to TUI experience
- [ ] create a game over state
- [ ] display hangman graphic
- [ ] support multiple difficulties
- [ ] document Setup, Use and Dev in `.md` files
- [ ] try inline docs?

How to build a binary so that i can run main as a command?

run main:
`go run cmd/main/main.go`

test all files in prject:
`go test ./...`
see percentage of code covered:
`go test -cover`
To build test covereage:
`go test -coverprofile=coverage.out`
To see it:
`go tool cover -html=coverage.out`

Flow (happens in session):
shell session -> calls New Game
New game is created, default states are initialized
list of words is read from a file
one is randinly selected -> this word goes is the one to be guessed
display current guess (empty but lenth equal word)
prompt user input
receive user input
validate
check
update game state
repeat until over
show game over
offer to replay or not
