# PathClip 
Copy the absolute path of a file to your clipboard

## Overview

`ptc` is a simple command-line tool designed to copy the absolute path of a given file or directory, directly to your clipboard. It aims to simplify the process of sharing or using file paths by making it a one-step operation.

## Features

- Cross-platform support (Windows, macOS, Linux)
- Copies the absolute path of a file to the clipboard
- Simple and easy-to-use command-line interface

## Installation

### From Source

1. Clone the repository:
    ```bash
    git clone https://github.com/marduc812/pathclip.git
    ```
2. Navigate to the project directory:
    ```bash
    cd pathclip
    ```
3. Build the project:
    ```bash
    go build
    ```

### Pre-compiled Binaries

Download the pre-compiled binaries for your operating system from the [Releases](https://github.com/marduc812/pathclip/releases) page.

## Usage

To copy the absolute path of a file to the clipboard, simply run:

```bash
ptc /path/to/file
```

Copy the content of a file:

```bash
ptc -c /path/to/file
```

## Contributing
Feel free to open issues or submit pull requests. All contributions are welcome!

## ChangeLog

[0.3] - 10/12/2024

- Added some color
- Added better errors for linux 
