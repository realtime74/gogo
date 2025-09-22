# gogo - a simple directory jumper

gogo is a command-line tool that allows you to quickly navigate to frequently used directories. It keeps track of your directory usage and provides an easy way to jump to them using a simple interface.

## Installation

1.  Install gogo

    ```bash
    go install github.com/realtime74/gogo@latest
    ```

2. Add this function to your shell configuration file (e.g., `~/.bashrc`, `~/.zshrc`):

    ```bash
    function gg() { cd $(gogo $1) }
    ```

## FUDs

FUDs (Frequently Used Directories) are stored in a file located at `~/.gogo`.

Each line in the file represents a directory,
with the most frequently used directories at the top.

## Usage

```bash
gogo [pattern|+|-|~]
```

- `pattern`: A substring to match against your frequently used directories. If no pattern is provided, gogo will
   jump to the home directory
- `+`: Add current directory to the list of frequently used directories
- `-`: Jump to the previous directory (cd -)

