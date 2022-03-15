# tmpl

![License](https://img.shields.io/github/license/Bananenpro/tmpl)
![Go Version](https://img.shields.io/github/go-mod/go-version/Bananenpro/tmpl)

## Description

A command line tool that lets you easily create new programming projects with support for many languages and frameworks.

## Supported languages/frameworks

- C
	- Manual
	- Makefile
	- CMake
- C#
	- Console
	- ASP.NET Core Empty
	- ASP.NET Core Web App
	- ASP.NET Core Web App (Model-View-Controller)
	- ASP.NET Core Web API
- Go
- Python
- Rust
	- Binary
	- Library

### Other

- Initialize git repository
- Create README.md file
- Create LICENSE file

## Usage

Create a new Go project in the current directory (must be empty):

```sh
tmpl new go
```

Help:

```sh
tmpl --help
```

## Setup

### Cloning the repo

```bash
git clone https://github.com/Bananenpro/tmpl.git
cd tmpl
```

### Requirements

- [Go](https://go.dev/) 1.17+

### Run

```sh
go run ./cmd
```

### Build

```sh
go build -o tmpl ./cmd/main.go
```

## License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

## Copyright

Copyright © 2022 Julian Hofmann
