# GoShell — A Tiny Terminal Written in Go

GoShell is a simple terminal app built with Go. It’s a lightweight shell that lets you run system commands and navigate around your file system — kind of like your regular terminal, just... homegrown.

---

## What It Does

- Shows your username, hostname, and current folder in the prompt (just like bash or zsh).
- Runs most commands you’d use in a normal shell (`ls`, `echo`, etc.).
- Supports `cd` to change directories. (home directory in "cd" not implemented)
- Supports `exit` to close the shell.
- Handles errors without crashing.

Here’s what the prompt looks like:

```
[alice@my-computer my-folder]$
```

---

## 🔧 How to Install & Run

First, make sure you have Go installed. If not, [download it here](https://golang.org/dl/).

1. Clone this repo:

```bash
git clone https://github.com/EqualNoches/go_shell.git
cd go_shell
```

2. Build it:

```bash
go build -o go_shell
```

3. Run it:

```bash
./go_shell
```

To test if it's really working, try typing `ls` or `echo Hello!`
---

## 🗺️ Commands You Can Use

This shell runs most system commands, but it also has a couple of built-in ones:

| Command | What It Does                        |
|---------|-------------------------------------|
| `cd`    | Change the current directory        |
| `exit`  | Quit the shell                      |

Anything else (like `ls`, `pwd`, `mkdir`, etc.) gets passed to your system's shell.

---

## 💡 Why I Built This

Mostly as a learning exercise. I wanted to get more comfortable with Go’s `os/exec`, work with input/output streams, and understand how shells handle command parsing.

That said, it turned out kind of cool. Maybe you’ll find it useful or want to build on top of it.

---

## 📌 Future Ideas

This is super basic right now. Here are some things I might (or you might!) add later:

- Command history with up/down arrows
- Tab completion
- Piping and redirection support
- Colored prompt

---

## 📝 License

MIT. Do whatever you want with it. If you improve it, I'd love to see what you build.

---

## 🙌 Thanks

If you made it this far — thanks for checking this out. If you have feedback or questions, feel free to open an issue or just fork and tinker.

Edward Díaz
