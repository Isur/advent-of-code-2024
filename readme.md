# Advent of Code

My solutions to the [Advent of Code](https://adventofcode.com).

I am using go.

# Progress

:white_check_mark:
:white_check_mark:
:white_check_mark:
:white_check_mark:
:white_check_mark:
:white_check_mark:
:white_check_mark:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:
:black_large_square:

- :black_large_square: - not done
- :black_square_button: - one star
- :white_check_mark: - two stars

# Run

Firstly for day you want to run create files `example` and `input`.

Example:
```bash
.
├── cmd
│   └── main.go
├── days
│   ├── Day01
│   │   ├── example
│   │   ├── input
│   │   └── run.go
│   ├── Day02
│   │   ├── example
│   │   ├── input
│   │   └── run.go
│   ├── Day03
│   │   ├── example
│   │   ├── input
│   │   └── run.go
```

Then to run:
```bash
go run ./cmd/main.go [day] [mode]
```

Where day is day you want to run and mode is `true` or `false` where `true` will use `input` and `false` will use `example` file.

By default it will run `example`.

Examples:

```bash
go run ./cmd/main.go 01 true
go run ./cmd/main.go 02
```
