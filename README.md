# ⚡ profile terminal

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Transform tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`transform` `data-processing` `cli` `golang`

---

## What is profile-terminal?

**profile-terminal** is a data transformation tool that converts, formats, and processes files between different formats.

## Features

- ✅ `svgBox()` — Svgbox
- ✅ `esc()` — Esc
- ✅ CLI flags and options
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/profile-terminal.git
cd profile-terminal

# Build
go build -o profile-terminal .

# Run
./profile-terminal -out -s
```

### Or directly with `go run`:
```bash
go run main.go -out -s
```

## Usage

```bash
# Basic usage
./profile-terminal -out -s

# With flags
./profile-terminal -out value -out -s
./profile-terminal -s value -out -s
```

### Example Output

```
$ ./profile-terminal -out -s
wrote %s (%d lines)\n
```

## Project Structure

```
profile-terminal/
  main.go          # Entry point (40 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
