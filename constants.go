package gen

const commonMakefile = `
lint:
	golangci-lint run

test:
	go test ./...
`

const MakefileE = `release-local:
	goreleaser release --snapshot --clean

release:
	goreleaser release
` + commonMakefile

const GitIgnore = "*.swp\nbuilds/\ntarget/"

const MakefileI = commonMakefile
