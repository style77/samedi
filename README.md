# Samedi

Samedi is a lightweight blogging tool written in Go that allows you to create and manage your blog(s) from the terminal.

It does not use any JavaScript and keeps the dependencies minimal, relying only on two dependencies: `blackfriday` for parsing Markdown and the SQLite driver.
Create new blog, new post and server the blog within your CLI.

## Features

- Create and manage multiple blogs from the terminal.
- Write and publish posts using Markdown.
- Simple and lightweight, no JavaScript involved.

## Installation

To install Samedi, you need to have Go installed on your system. Then, you can install it using `go get`:

```bash
go install github.com/style77/samedi
```

This will download the Samedi project and install it into your $GOPATH/bin directory.

## Usage

After installing Samedi, you can start using it from the terminal:

```
samedi help
```

## Contributing

Contributions to Samedi are welcome! If you find any bugs or have ideas for new features, feel free to open an issue or submit a pull request on GitHub.
License

## License

Samedi is licensed under the MIT License. See the LICENSE file for details.