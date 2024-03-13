> The development of the package was discontinued, see [mittelweg](https://github.com/bernardoforcillo/mittelweg) and [diskursus](https://github.com/bernardoforcillo/diskursus)

# A Golang SDK for managing Midjourney through Discord API.

The `midjourney-go` package provides a Golang SDK for managing Midjourney through Discord API. It is in development.

## Installation

```bash
go get github.com/bernardoforcillo/midjourney-go
```

## Usage

For use in your project, add the `midjourney-go` package to your project. As follows:

```go
import "github.com/bernardoforcillo/midjourney-go/midjourney"
```

### How to initialize a client

```go
client := midjourney.NewMidjourneyClient("your token", "channel id")

```

### How to generate an image

The `Imagine` command is used to generate an image. It takes two arguments: a string that is the prompt for the image, and a boolean that force the execution to wait until the image is generated.

```go
generatedImage, err := client.Imagine("prompt", waitUntilGenerated)
if err != nil {
    log.Fatalf("call client.Imagine failed, err: %+v", err)
}
```

### How to upscale an image

The `Upscale` command is used to upscale an image. It takes two arguments: the index of the image to upscale, the index must be between 0 (included) and 4(excluded). and a boolean that force the execution to wait until the image is generated.

```go
upscaledImage, err := generatedImage.Upscale(index, true)
if err != nil {
	log.Fatalf("call.Upscale failed, err: %+v", err)
}
```

`SearchMesssageWithContent` and `SearchMesssageByPrompt` are two utility functions
that can be used to search for messages in a channel.
## License

This project is licensed under the MIT license. See the [license.md](license.md) file for more details.

---

midjourney-go is a Go client library for accessing the [Midjourney-Bot API](https://midjourney.com/).

It is designed very simply and is lightweight with no additional logic. It belongs to a very low-level library, so you can use it to do anything.

If you want to build your own midjourney apiserver, take a look at this [midjourney-apiserver](https://github.com/hongliang5316/midjourney-apiserver).

[List](#list)

## Installation ##

midjourney-go is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/hongliang5316/midjourney-go@v0.0.1
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/hongliang5316/midjourney-go/midjourney"
```

and run `go mod tidy` without parameters.

[List](#list)

## Simple Usage ##

Construct a new `Midjourney client`, then use the various commands on the client to
access different parts of the `Midjourney-Bot API`. For example:

```go
client := midjourney.NewClient(&midjourney.Config{
    UserToken: "your token",
})

// imagine
if err := client.Imagine(context.Background(), &midjourney.ImagineRequest{
    GuildID: "",
    ChannelID: "",
    Prompt: "",
}); err != nil {
    log.Fatalf("Call client.Imagine failed, err: %+v", err)
}
```

[List](#list)

## Features ##

Currently, only some [commands](https://docs.midjourney.com/docs/command-list) have been implemented. All commands will be implemented in the future. If you are interested, please submit a pull request or issues.

- [x] /imagine

- [x] /upscale

- [x] /variation

- [x] /describe

- [ ] /blend

- [ ] /reroll

[List](#list)
