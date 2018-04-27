generate-plugin
===============
`generate-plugin` is a [godel](https://github.com/palantir/godel) plugin for [`go-generate`](https://github.com/palantir/go-generate). It provides a task that runs "go generate" based on project configuration and a verification task that verifies that the current state of the project matches what is expected. 

Tasks
-----
* `generate`: runs the "go generate" task based on the plugin configuration. 

Verify
------
When run as part of the `verify` task, if `apply=true`, then the `generate` task is run. If `apply=false`, then `generate --verify` is run, which verifies that the running "generate" does not modify the content of the target output files. Note that, even when "apply=false", the task may change state on disk (because there is no way to control what the individual "generate" tasks actually do).
