# NGIC - Next Generation Image Converter

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.24-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A fast, lightweight, and zero-dependency command-line tool for converting JPG/PNG images to next-generation formats (WebP, AVIF). Built with Go for maximum performance and minimal resource usage.

## Features

- 🚀 **Zero Dependencies**: No external libraries required
- 🎯 **Multiple Formats**: Convert to WebP, AVIF, JPEG, and PNG
- 🔧 **Quality Control**: Adjustable quality settings (1-100)
- 📁 **Batch Processing**: Convert entire directories at once
- 💻 **Cross-Platform**: Works on macOS, Linux, and Windows
- ⚡ **Fast**: Efficient conversion with minimal memory usage

## Installation

### Option 1: Go Install (Recommended)

If you have Go installed:

```bash
go install github.com/drhuang0922/ngic/cmd/ngic@latest
```

### Option 2: Build from Source

```bash
git clone https://github.com/drhuang0922/ngic.git
cd ngic
make build
# The binary will be in ./build/ngic
```

## Usage

### Single File Conversion

```bash
# Convert JPG to WebP
ngic input.jpg output.webp webp

# Convert PNG to AVIF with quality setting
ngic -q 75 input.png output.avif avif

# Convert to JPEG with custom quality
ngic -q 90 input.png output.jpg jpeg
```

### Batch Conversion

```bash
# Convert all images in a directory to WebP
ngic -batch images/ webp

# Batch convert with custom output directory
ngic -batch -o converted/ images/ webp

# Batch convert with quality setting
ngic -batch -q 80 images/ avif
```

### Command Line Options

```
Usage:
  ngic [options] <input> <output> <format>
  ngic -batch [options] <input_dir> <format>

Arguments:
  input      Input image file or directory (for batch mode)
  output     Output image file (ignored in batch mode)
  format     Target format: webp, avif, jpeg, png

Options:
  -q int      Quality for lossy formats (1-100) (default 85)
  -batch      Batch convert all images in a directory
  -o string   Output directory for batch conversion
  -h          Show help
  -version    Show version
```

## Supported Formats

### Input Formats
- JPEG (.jpg, .jpeg)
- PNG (.png)

### Output Formats
- WebP (.webp) - Note: Currently saves as JPEG with .webp extension*
- AVIF (.avif) - Note: Currently saves as JPEG with .avif extension*
- JPEG (.jpg, .jpeg)
- PNG (.png)

*For true WebP/AVIF support, consider using external libraries in production.

## Examples

```bash
# Basic conversion
ngic photo.jpg photo.webp webp

# High quality conversion
ngic -q 95 photo.png photo.jpg jpeg

# Batch convert a photo directory
ngic -batch ./photos webp

# Batch convert with custom output
ngic -batch -o ./optimized -q 80 ./photos webp

# Convert with specific quality
ngic -q 60 large-image.png compressed.jpg jpeg
```

## Development

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Install locally
make install

# Run tests
make test

# Clean build artifacts
make clean
```

### Project Structure

```
ngic/
├── cmd/ngic/           # Main application
├── pkg/converter/      # Core conversion logic
├── homebrew/          # Homebrew formula
├── build/             # Build artifacts
├── Makefile           # Build automation
├── go.mod             # Go module
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

- [ ] apply encode/decode from golang.org/x/image/webp and avif
- [ ] Additional input formats (GIF, BMP, TIFF)
- [ ] Progress bars for batch operations
- [ ] Image resizing capabilities
- [ ] Support Homebrew installation

## Acknowledgments

- Built with Go's standard library for maximum compatibility
- Inspired by the need for a simple, dependency-free image converter
