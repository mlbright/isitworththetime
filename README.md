# isitworththetime

> Is creating an automation worth it?

## What

This is a toy codebase meant to exercise some Golang and GitHub Actions features and capabilities.

How much time can you sink into an automation before you are spending more time than you save?

Assumptions:

- work day is 8 hours
- work week is 5 days

## Why

The inspiration for this is [the XKCD comic][isitworththetime].

The purpose of this repository to develop an example Golang project that exercises all the relevant GitHub features for the project.

It includes:

- a library (`pkg` directory)
- a CLI (`cmd` directory)
- a README.md
- tests
- versioned releases
- a web assembly web page

It does not include:

- a GitHub wiki
- a GitHub project

## How

### Install the CLI

Copy the binary executable into your system PATH.
Then run:

```bash
./isitworththetime --help
```

### Use the library

```bash
go get -u https://github.com/mlbright/isitworththetime/pkg
```

In your code:

```go
import (
	automation "github.com/mlbright/isitworththetime/pkg"
)
```

### Usage

```bash
isitworththetime --saved 30s --interval 672h --frequency 1
```

### Build the executable file

```bash
make
```

[isitworththetime]: https://xkcd.com/1205/
