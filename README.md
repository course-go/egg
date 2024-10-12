# EGG

EGG is a _Even Greater Grep_ command line application.
This project serves as a template for the first homework assignment.
To learn more about the homework assignments in general, visit
the [homework](https://github.com/course-go/homework) repository.

## Assignment

Throughout this homework assignment you will create a clone of the
[grep](https://man7.org/linux/man-pages/man1/grep.1.html) command line tool.

### Specification

The basic syntax is as follows:

```shell
egg pattern [FILES]
```

where **pattern** represents a regular expression as supported by the
[regexp](https://pkg.go.dev/regexp) package and **files** represents a path
or list of paths to a single or multiple files. If no files are
provided, the grep will instead read data from the standard input.

The example usage is as follows:

```shell
egg '\[[0-9]+\]' ./my/cool/file.txt /etc/random/config.txt
```

The application will then read all the specified files and searches for lines
that match the regular expression. For each line that matches, the
application will print the matched line together with the line number.

The sample output:

```text
File: ./my/cool/file.txt
    12: coming up with examples is hard [12]. It is not expected
    25: “The computer was born to solve problems that did not exist before.” [23]
```

When an invalid input is provided to the application or the application
fails, it must print an error message to the standard error output and exit
with non-zero status code.

If you are unsure about some behaviour take the tests as the source of truth.

## Requirements

The CLI application has to support all the functionality previously
described. Stress is also given on writing idiomatic code and properly
handling resources and errors.

## Motivation

The main goal of this homework is to practice basic control flow, error
handling and interaction with several Go packages and their documentations.

## Packages

Some of the Go packages worth looking into include:

- [os](https://pkg.go.dev/os) for interacting with the operation system
- [io](https://pkg.go.dev/io) and [bufio](https://pkg.go.dev/bufio) for interacting
with input and outputs
- [regexp](https://pkg.go.dev/regexp) for working with the regular expressions
