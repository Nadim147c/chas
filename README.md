# chas (char + has)

`chas` is a lightweight CLI utility that reads lines from standard input and filters
them based on whether they contain a specific set of characters.

## Installation

```bash
go install github.com/Nadim147c/chas@latest

```

## Usage

Pass the characters you are looking for as arguments, and pipe your text into `chas`.

```bash
# Basic usage
chas abc < file.txt

# Using multiple arguments (joined automatically)
echo "hello world" | chas hlo

```

### How it works

`chas` performs an **AND** operation on the provided characters. A line will only be
printed to `stdout` if **all** the characters specified are present somewhere within
that line.

| Input    | Command    | Output      | Reason                     |
| -------- | ---------- | ----------- | -------------------------- |
| `apple`  | `chas ap`  | `apple`     | Contains both 'a' and 'p'  |
| `apple`  | `chas az`  | _(nothing)_ | Missing 'z'                |
| `banana` | `chas nna` | `banana`    | Contains 'n', 'n', and 'a' |

## License

Licensed under [GPL-3.0](./LICENSE)!
