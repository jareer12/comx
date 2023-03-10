# ComX - A C/C++ Project Manager CLI

ComX is a lightweight and user-friendly command-line interface (CLI) tool for managing C/C++ projects. With ComX, you can easily create and configure new projects, manage source code files, and compile your projects with just a few simple commands. Its minimalistic design and ease-of-use make it a great choice for developers who want a simple yet powerful tool for managing their C/C++ projects. Save time and streamline your workflow with ComX.

## Features

- Lightweight
- Easy to use
- Portable

## Installation

ComX is currently available as a standalone binary. To install it on your system, follow these steps:

### Windows

1. Download the latest release from the [GitHub repository](https://github.com/jareer12/comx/releases).
2. Extract the binary file from the downloaded archive.
3. Copy the binary to a location in your system's PATH, such as `/usr/local/bin`. Rename it to `comx` for convenient use.

### Linux

Run the following script in your terminal.

```sh
bash https://raw.githubusercontent.com/jareer12/comx/main/install.sh
```

## Usage

To create a new C project with ComX, simply run the following command:

### Create New Project

```sh
## comx init <project-name>
comx init app
```

### Building Project

```sh
## comx build -o ./binary [--args]
comx build
```

## Building From Source

Clone this repository and run the `./scripts/build.sh` file, do not use `sudo` as it might not work. If an error is thrown anyways, run `sudo chmod -R 777 ./*`.
