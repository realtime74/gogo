# gogo - a simple directory jumper

gogo is a command-line tool that allows you to quickly navigate to frequently used directories. It keeps track of your directory usage and provides an easy way to jump to them using a simple interface.

## Usage

```bash
gogo [pattern|+|-|~]
```

- `pattern`: A substring to match against your frequently used directories. If no pattern is provided, gogo will
   jump to the home directory
- `+`: Add current directory to the list of frequently used directories
- `-`: Jump to the previous directory (cd -)
