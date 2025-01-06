# 7solutions Backend Challenges (Golang)

My attempt on 7solutions backend software engineer challenges (Golang),

For more detail, please check https://github.com/7-solutions/backend-challenge/

## Get started

Clone the repository and grab all dependencies

```sh
git clone git@github.com:ball6847/7solutions-backend-challenges.git
cd 7solutions-backend-challenges
go get -d ./...
```
## How to use the command line interface

All the challenges can be started using command line interface. Try running the cli with `--help` flag, you will see all available commands

```
$ go run . --help
Usage:
  7solutions [command]

Available Commands:
  challenge1  Run and see the result for challenge1
  challenge2  Run and see the result for challenge2
  challenge3  Start api server for challenge3
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for 7solutions

Use "7solutions [command] --help" for more information about a command.
```

### Challenge 1

```
$ go run . challenge1
The answer for challenge #1 is:  7273
```

- The input dataset file [hard.json](hard.json) has been copied from the source repository
- The challenge is solved using **Bottom-up Dynamic Programming**, check source code for more detail [cmd/challenge1/challenge1.go](cmd/challenge1/challenge1.go)

### Challenge 2

By running `challenge2` subcommand, you will be prompt for input and the result will be returned after hitting `ENTER`.

```
$ go run . challenge2
Please enter your encoded input:LLRR=
Total solution: 2892
Best solution: 210122
```

- As the challenge could produce a lot of possible solutions, the optimal solution will be picked automatically based on the rules mentioned in the source repository.
- The challenge is solved using **Backtrack Algorithm**, check source code for more detail - [cmd/challenge2/challenge2.go](cmd/challenge2/challenge2.go)

For list of all possible solutions, the `--dump-solutions` is provided for debugging purpose.

```
$ go run . challenge2 --dump-solutions
Please enter your encoded input:LLRR=
Total solution: 2892
Best solution: 210122
All solutions has been dumped to solutions.txt
```

Then check out `solutions.txt` for the possible solutions

### Challenge 3

Challenge 3 is rest api implemented using Gin.

```
$ go run . challenge3
```

The server will be available at `http://localhost:5555/beef/summary`

## Running the tests

```
$ go test ./...
