# Profile-Terminal

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-00ADD8?style=for-the-badge)
![CI](https://img.shields.io/badge/CI-GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)
![CodeQL](https://img.shields.io/badge/CodeQL-Security-00ADD8?style=for-the-badge)
![Version](https://img.shields.io/badge/Version-v1.0.0-00ADD8?style=for-the-badge)

> Turn your GitHub profile into an interactive ASCII terminal with dynamic SVG

`github-profile` `terminal` `svg` `interactive` `serverless` `developer-tools` `golang`

---

## What is it?

**Profile-Terminal** is Replace your static GitHub profile with a living terminal. Visitors type commands, view stats, explore projects - all rendered as dynamic SVG.

## Why should you care?

- **Fast** - Compiled Go binary, zero overhead
- **Secure** - CodeQL analysis + Dependabot
- **Offline-first** - Works without internet
- **Lightweight** - Single binary deployment
- **Developer-friendly** - Clean CLI with docs

---

## Features

- Interactive terminal UI in README
- Dynamic SVG rendering
- Command processing (help, skills, projects)
- Real-time GitHub stats
- Visitor counter
- Animated typing effect
- Color themes (matrix, cyberpunk, retro)
- Serverless backend (Cloudflare Workers)
- Custom command aliases
- Easter eggs

---

## Quick Start

### Prerequisites
- Go 1.21 or higher

### Install from source
```bash
git clone https://github.com/AetherCodeHQ/profile-terminal.git
cd profile-terminal
go build -o profile-terminal .
```

### Run
```bash
./profile-terminal --help
```

---

## Usage

./termctl generate --theme matrix --commands help,skills,projects  OR  ./termctl deploy --provider cloudflare

---

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--path` | Target directory | `.` |
| `--format` | Output format (json, text) | `text` |
| `--output` | Output filename | `stdout` |
| `--verbose` | Enable verbose output | `false` |

---

## Development

```bash
git clone https://github.com/AetherCodeHQ/profile-terminal.git
cd profile-terminal
go build -o profile-terminal .
go test ./...
golangci-lint run
```

---

## Contributing

Contributions welcome! See [CONTRIBUTING.md](CONTRIBUTING.md).

## Security

Report to: aethercode.core@gmail.com | See [SECURITY.md](SECURITY.md)

## License

MIT License - see [LICENSE](LICENSE)

---

<p align="center">
  Built with love by <a href="https://github.com/AetherCodeHQ">AetherCode</a> | <a href="https://github.com/AetherCode-Core">AetherCode Core</a>
</p>
