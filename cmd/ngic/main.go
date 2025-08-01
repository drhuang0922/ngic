package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/drhuang0922/ngic/pkg/converter"
)

const version = "1.0.0"

func main() {
	var (
		quality   = flag.Int("q", 85, "Quality for lossy formats (1-100)")
		batch     = flag.Bool("batch", false, "Batch convert all images in a directory")
		help      = flag.Bool("h", false, "Show help")
		showVer   = flag.Bool("version", false, "Show version")
		outputDir = flag.String("o", "", "Output directory for batch conversion")
	)

	flag.Usage = func() {
		fmt.Printf("ngic - Next Generation Image Converter v%s\n\n", version)
		fmt.Println("Usage:")
		fmt.Println("  ngic [options] <input> <output> <format>")
		fmt.Println("  ngic -batch [options] <input_dir> <format>")
		fmt.Println("")
		fmt.Println("Arguments:")
		fmt.Println("  input      Input image file or directory (for batch mode)")
		fmt.Println("  output     Output image file (ignored in batch mode)")
		fmt.Println("  format     Target format: webp, avif, jpeg, png")
		fmt.Println("")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Println("  ngic input.jpg output.webp webp")
		fmt.Println("  ngic -q 75 input.png output.jpg jpeg")
		fmt.Println("  ngic -batch images/ webp")
		fmt.Println("  ngic -batch -o converted/ images/ webp")
	}

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *showVer {
		fmt.Printf("ngic version %s\n", version)
		return
	}

	args := flag.Args()

	// Validate quality
	if *quality < 1 || *quality > 100 {
		fmt.Fprintf(os.Stderr, "Error: Quality must be between 1 and 100\n")
		os.Exit(1)
	}

	conv := converter.NewImageConverter()
	conv.SetQuality(*quality)

	if *batch {
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Error: Batch mode requires input directory and format\n")
			flag.Usage()
			os.Exit(1)
		}

		inputDir := args[0]
		format := args[1]

		// Validate input directory
		if !isDirectory(inputDir) {
			fmt.Fprintf(os.Stderr, "Error: Input path '%s' is not a directory or doesn't exist\n", inputDir)
			os.Exit(1)
		}

		// Validate format
		if !isValidFormat(format) {
			fmt.Fprintf(os.Stderr, "Error: Unsupported format '%s'. Supported: webp, avif, jpeg, png\n", format)
			os.Exit(1)
		}

		// Determine output directory
		var outDir string
		if *outputDir != "" {
			outDir = *outputDir
		} else {
			outDir = filepath.Join(inputDir, "converted")
		}

		fmt.Printf("Batch converting images from %s to %s (format: %s, quality: %d)\n",
			inputDir, outDir, format, *quality)

		if err := conv.BatchConvert(inputDir, outDir, format); err != nil {
			fmt.Fprintf(os.Stderr, "Batch conversion failed: %v\n", err)
			os.Exit(1)
		}
	} else {
		if len(args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: Single file mode requires input file, output file, and format\n")
			flag.Usage()
			os.Exit(1)
		}

		inputPath := args[0]
		outputPath := args[1]
		format := args[2]

		// Validate input file
		if !fileExists(inputPath) {
			fmt.Fprintf(os.Stderr, "Error: Input file '%s' doesn't exist\n", inputPath)
			os.Exit(1)
		}

		// Validate input file format
		if !isValidInputFile(inputPath) {
			fmt.Fprintf(os.Stderr, "Error: Input file '%s' is not a supported image format (jpg, jpeg, png)\n", inputPath)
			os.Exit(1)
		}

		// Validate format
		if !isValidFormat(format) {
			fmt.Fprintf(os.Stderr, "Error: Unsupported format '%s'. Supported: webp, avif, jpeg, png\n", format)
			os.Exit(1)
		}

		// Show input image info
		if width, height, inputFormat, err := converter.GetImageInfo(inputPath); err == nil {
			fmt.Printf("Input: %s (%dx%d, %s)\n", inputPath, width, height, inputFormat)
		}

		fmt.Printf("Converting %s to %s (format: %s, quality: %d)\n",
			inputPath, outputPath, format, *quality)

		if err := conv.ConvertImage(inputPath, outputPath, format); err != nil {
			fmt.Fprintf(os.Stderr, "Conversion failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully converted to %s\n", outputPath)
	}
}

// Helper functions
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func isValidInputFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

func isValidFormat(format string) bool {
	format = strings.ToLower(format)
	return format == "webp" || format == "avif" || format == "jpeg" || format == "jpg" || format == "png"
}
