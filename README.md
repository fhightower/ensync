# EnSync

Highlighting differences between env files (e.g. `.env.example` and `.env`).

## Setup

1. Clone the repo
2. Run `go get` in the repo dir
3. Run `go run main.go <path>`

## Usage

There are two use-cases for this tool:

1. Comparing two env files in the current dir
2. Comparing all env file pairs in all subdirs of the current dir

If the script files both a `.env` and `.env.example` in the current dir, it will compare them and stop.

If there either one of these files is missing from the current dir, it will look for them in all subdirs of the current dir.

