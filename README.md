# Lacking Template

This is a Go template project for the [lacking](https://github.com/mokiat/lacking) game engine.

## Usage

Install [gonew](https://pkg.go.dev/golang.org/x/tools/cmd/gonew), if you don't have it already:

```sh
go install golang.org/x/tools/cmd/gonew@latest
```

Create a new project from this template:

```sh
gonew github.com/mokiat/lacking-template@latest example.com/your/namespace projectdir
```

Then, from within the project directory, prepare the assets and run the game using [Task](https://taskfile.dev):

```sh
task pack
task run
```

Run `task --list` to see all available tasks.

## References

- [lacking](https://github.com/mokiat/lacking) - the game engine
- [gonew](https://go.dev/blog/gonew) - starting new projects from templates
- [Task](https://taskfile.dev) - the task runner used by this project
