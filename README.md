
# gitignore - A simple .gitignore generator

A simple cli tool written in go to generates gitignore files for different projects.

## Installation

### Download

You can download the latest release [here](https://github.com/ddddami/gitignore/releases). Extract the zip file, and run the executable, like;

```sh
cd gitignore && ./gitignore -help
```

You can add this to your PATH too.

```sh
# Linux/macOS
sudo mv gitignore /usr/local/bin/
```

### From Source

1. Clone the repository:

```sh
git clone https://github.com/ddddami/gitignore
cd gitignore
```

2. Build the binary.

```sh
go build -o gitignore ./cmd/gitignore
```

3. Move the binary to your PATH

```sh
# Linux/macOS
sudo mv gitignore /usr/local/bin/
```

## Usage

```sh
gitignore [template-name]
```

For example:

```
gitignore node   # Generates a Node.js .gitignore
gitignore python # Generates a Python .gitignore
```

Running `gitignore` without arguments will display available templates.

## Adding New Templates

This is a very simple tool, I added some templates from [github/gitignore](https://github.com/gitignore)

### To add new templates

Create a new file in the templates directory with the naming convention [template-name].gitignore and add the content

Rebuild the application

## Extending the Tool

Some ideas for extending this tool:

- Add support for combining multiple templates [gitignore node,python gitignore node python]
- Add support for adding templates to custom dirs
- Add a flag to append to an existing .gitignore instead of creating/trying to create a new one
- Fetch templates remotely (from GitHub)
