package fail

fail

/*
This is a non-compiling file that has been added to explicitly ensure that CI fails.
It also contains the command that caused the failure and its output.
Remove this file if debugging locally.

go mod operation failed. This may mean that there are legitimate dependency issues with the "go.mod" definition in the repository and the updates performed by the gomod check. This branch can be cloned locally to debug the issue.

Command that caused error:
./godelw exec -- go get -d github.com/palantir/go-generate@upgrade github.com/palantir/godel/v2@upgrade github.com/palantir/pkg/cobracli@upgrade github.com/palantir/pkg/matcher@upgrade github.com/palantir/pkg/pkgpath@upgrade github.com/palantir/pkg/specdir@upgrade github.com/spf13/cobra@upgrade github.com/spf13/pflag@upgrade github.com/stretchr/testify@upgrade golang.org/x/tools@upgrade gopkg.in/yaml.v2@upgrade gopkg.in/yaml.v3@upgrade

Output:
go: -d flag is deprecated. -d=true is a no-op
go: downloading github.com/palantir/go-generate v1.37.0
go: github.com/palantir/go-generate@upgrade: github.com/palantir/go-generate@v1.37.0: verifying module: github.com/palantir/go-generate@v1.37.0: reading https://sum.golang.org/lookup/github.com/palantir/go-generate@v1.37.0: 404 Not Found
	server response:
	not found: github.com/palantir/go-generate@v1.37.0: invalid version: git ls-remote -q origin in /tmp/gopath/pkg/mod/cache/vcs/a66cd17bb12d3e3d357df15a6080fb5a6ae995210ea0bfffd4ceff11920eca15: exit status 128:
		fatal: unable to access 'https://github.com/palantir/go-generate/': Failed to connect to github.com port 443: Connection refused
Error: exit status 1

*/
