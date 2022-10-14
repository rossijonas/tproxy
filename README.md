<h1 align="center">tproxy</h1>

*<p align="center">A simple TCP proxy built with Go.</p>*

## About

`tproxy` do what????

### Status:

[![Actions Status](https://github.com/rossijonas/tproxy/workflows/Build/badge.svg)](https://github.com/rossijonas/tproxy/actions)

### Features:

- Cross platform:  Linux / Macos / Windows.

- What Features????

_(This is an exercise from the book "Black Hat Go", although it may have some differences from the original exercise.)_

### **🚫⚠️  Warning!**

**This repository's content must NOT be used for any activity considered illegal in any country in the world! Stay away from this repository if you are not sure what you are doing or if you will not use this for only legal activities! I won't take responsibility for any consequences of your use of this content!**

## Installation

### Requirements:

- [Go](https://go.dev/) version 1.18.6 (or above)

### How to install:

- Run: 

  ```
  $ go install github.com/rossijonas/tproxy@latest
  ```

## Usage

### Options:

```
$ wc -h
Usage of wc:
  -b    Count bytes
  -l    Count lines
```

### Examples:

Pipe text to `wc` command, it returns the number of words by default.

#### Count words:

```
$ echo "two words" | ./wc
2 # result
```

#### Count lines using `-l` flag:

```
$ cat main.go | ./wc -l
32 # result
```

## Backlog

- Establish a connection to a remote listener via `net.Dial(network, address string)`.

- Initialize a `Cmd` via `exec.Command(name string, arg ...string)`.

- Redirect `Stdin` and `Stdout` properties to utilize the `net.Conn` object.

- Run the command


