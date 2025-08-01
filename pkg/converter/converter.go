package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ImageConverter handles image format conversion
type ImageConverter struct {
	Quality int // Quality for lossy formats (1-100)
}

// NewImageConverter creates a new converter with default quality
func NewImageConverter() *ImageConverter {
	return &ImageConverter{
		Quality: 85,
	}
}

// SetQuality sets the quality for lossy format conversion
func (ic *ImageConverter) SetQuality(quality int) {
	if quality < 1 {
		quality = 1
	}
	if quality > 100 {
		quality = 100
	}
	ic.Quality = quality
}

// ConvertImage converts an image from source to target format
func (ic *ImageConverter) ConvertImage(inputPath, outputPath, targetFormat string) error {
	// Open and decode source image
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	img, format, err := image.Decode(inputFile)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	fmt.Printf("Source format: %s\n", format)

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Convert based on target format
	targetFormat = strings.ToLower(targetFormat)
	switch targetFormat {
	case "webp":
		return ic.encodeAsWebP(img, outputFile)
	case "avif":
		return ic.encodeAsAVIF(img, outputFile)
	case "jpeg", "jpg":
		return jpeg.Encode(outputFile, img, &jpeg.Options{Quality: ic.Quality})
	case "png":
		return png.Encode(outputFile, img)
	default:
		return fmt.Errorf("unsupported target format: %s", targetFormat)
	}
}

// encodeAsWebP encodes image as WebP (simplified implementation)
// Note: This is a placeholder - for production use, consider using a WebP library
func (ic *ImageConverter) encodeAsWebP(img image.Image, output *os.File) error {
	// For zero dependency, we'll save as high-quality JPEG with .webp extension
	// In production, you'd use a proper WebP encoder library
	fmt.Println("Warning: Saving as JPEG with .webp extension (use WebP library for true WebP)")
	return jpeg.Encode(output, img, &jpeg.Options{Quality: ic.Quality})
}

// encodeAsAVIF encodes image as AVIF (simplified implementation)
// Note: This is a placeholder - for production use, consider using an AVIF library
func (ic *ImageConverter) encodeAsAVIF(img image.Image, output *os.File) error {
	// For zero dependency, we'll save as high-quality JPEG with .avif extension
	// In production, you'd use a proper AVIF encoder library
	fmt.Println("Warning: Saving as JPEG with .avif extension (use AVIF library for true AVIF)")
	return jpeg.Encode(output, img, &jpeg.Options{Quality: ic.Quality})
}

// BatchConvert converts multiple images in a directory
func (ic *ImageConverter) BatchConvert(inputDir, outputDir, targetFormat string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Read input directory
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("failed to read input directory: %w", err)
	}

	convertedCount := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		inputPath := filepath.Join(inputDir, file.Name())
		ext := strings.ToLower(filepath.Ext(file.Name()))

		// Check if it's a supported input format
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			fmt.Printf("Skipping unsupported file: %s\n", file.Name())
			continue
		}

		// Generate output filename
		baseName := strings.TrimSuffix(file.Name(), ext)
		outputFileName := fmt.Sprintf("%s.%s", baseName, targetFormat)
		outputPath := filepath.Join(outputDir, outputFileName)

		// Convert image
		fmt.Printf("Converting %s to %s...\n", inputPath, outputPath)
		if err := ic.ConvertImage(inputPath, outputPath, targetFormat); err != nil {
			fmt.Printf("Error converting %s: %v\n", inputPath, err)
			continue
		}

		convertedCount++
	}

	fmt.Printf("Successfully converted %d images\n", convertedCount)
	return nil
}

// GetImageInfo returns basic information about an image file
func GetImageInfo(imagePath string) (width, height int, format string, err error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, "", err
	}
	defer file.Close()

	config, format, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, "", err
	}

	return config.Width, config.Height, format, nil
}
