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

Create a new project in the current directory (must be empty):

```sh
tmpl
```

Use `--help` for a comprehensive list of options.

## Installation

### Windows

1. Open the Start menu
2. Search for `powershell`
3. Hit `Run as Administrator`
4. Paste the following commands and hit enter:

#### Install

```powershell
Invoke-WebRequest -Uri "https://github.com/Bananenpro/tmpl/releases/latest/download/tmpl-windows-amd64.zip" -OutFile "C:\Program Files\tmpl.zip"
Expand-Archive -LiteralPath "C:\Program Files\tmpl.zip" -DestinationPath "C:\Program Files\tmpl"
rm "C:\Program Files\tmpl.zip"
Set-ItemProperty -Path 'Registry::HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\Session Manager\Environment' -Name PATH -Value "$((Get-ItemProperty -Path 'Registry::HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\Session Manager\Environment' -Name PATH).path);C:\Program Files\tmpl"
```

**IMPORTANT:** Please reboot for the installation to take effect.

#### Update

```powershell
rm -r -fo "C:\Program Files\tmpl"
Invoke-WebRequest -Uri "https://github.com/Bananenpro/tmpl/releases/latest/download/tmpl-windows-amd64.zip" -OutFile "C:\Program Files\tmpl.zip"
Expand-Archive -LiteralPath "C:\Program Files\tmpl.zip" -DestinationPath "C:\Program Files\tmpl"
rm "C:\Program Files\tmpl.zip"
```

### macOS

Open the Terminal application, paste the command for your architecture and hit enter.

To update, simply run the command again.

#### x86_64

```sh
curl -L https://github.com/Bananenpro/tmpl/releases/latest/download/tmpl-darwin-amd64.tar.gz | tar -xz tmpl && sudo mv tmpl /usr/local/bin
```

#### ARM64

```sh
curl -L https://github.com/Bananenpro/tmpl/releases/latest/download/tmpl-darwin-arm64.tar.gz | tar -xz tmpl && sudo mv tmpl /usr/local/bin
```

### Linux

Open a terminal, paste the command for your architecture and hit enter.

To update, simply run the command again.

#### x86_64

```sh
curl -L https://github.com/Bananenpro/tmpl/releases/latest/download/tmpl-linux-amd64.tar.gz | tar -xz tmpl && sudo mv tmpl /usr/local/bin
```

#### ARM64

```sh
curl -L https://github.com/Bananenpro/tmpl/releases/latest/download/tmpl-linux-arm64.tar.gz | tar -xz tmpl && sudo mv tmpl /usr/local/bin
```

### Other

You can download a prebuilt binary file for your operating system on the [releases](https://github.com/Bananenpro/tmpl/releases) page.

You might need to make the file executable before running it.

### Compiling from source

#### Prerequisites

- [Go](https://go.dev/) 1.17+

```sh
git clone https://github.com/Bananenpro/tmpl.git
cd tmpl
go build .
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
