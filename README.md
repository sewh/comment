comment
=======

comment is a simple tool that adds or removes comments. It reads from stdin and outputs to stdout.

## Installation

```bash
go get github.com/sewh/comment
go install github.com/sewh/comment

# Make sure $(go env GOPATH)/bin is in your path
```

## Usage

To add a comment: `comment + comment_string` where `comment_string` is the comment string you wish to insert. For example, to add a '//' comment at the start of every line: `comment + '//'`.

To remove a comment: `comment - comment_regex`. For example, to remove a '//' at the start of every line: `comment - '^//'`.
