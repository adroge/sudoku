# Sudoku Solver

This is a Sudoku solver for any sized Sudoku.

Not sure what Sudoku is? Check here: [What is Sudoku?](https://duckduckgo.com/?q=what+is+sudoku)

## How to use

- Install Go - [here](https://go.dev/doc/install)
- Clone or download this repo - [docs](https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository)
- run the app
- Profit

## Example

```sh
go run ./sudoku.go
[1 2 3 4 5 6 7 8 9]
[4 5 6 7 8 9 1 2 3]
[7 8 9 1 2 3 4 5 6]
[2 1 4 3 6 5 8 9 7]
[3 6 5 8 9 7 2 1 4]
[8 9 7 2 1 4 3 6 5]
[5 3 1 6 4 2 9 7 8]
[6 4 2 9 7 8 5 3 1]
[9 7 8 5 3 1 6 4 2]
Show another solution [y/n]?
```

You can comment and uncomment the sections in the table array, set their
values, have fun, make this better, etc....

If you try solving a 25x25 Sudoku, it is going to take a long time. No
guarantees here.
