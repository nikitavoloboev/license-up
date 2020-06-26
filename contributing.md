# Contributing

Thank you for taking the time to contribute! ♥️ You can:

- Submit [bug reports or feature requests](../../issues/new/choose). Contribute to discussions. Fix [open issues](../../issues).
- Improve docs, the code and more! Any idea is welcome.

## Run project

To run the project:

1. Clone the repo
2. `go get` dependencies used. I
3. `go run .` to compile and run.

I like to use [watchexec](https://github.com/watchexec/watchexec) to recompile & run on go file changes. Using this command:

`watchexec --exts go "echo -- && go run ."`
