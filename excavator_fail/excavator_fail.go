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
go: github.com/palantir/go-generate@v1.37.0: Get "https://proxy.golang.org/github.com/palantir/go-generate/@v/v1.37.0.mod": dial tcp 142.250.65.177:443: i/o timeout
Error: exit status 1

*/
