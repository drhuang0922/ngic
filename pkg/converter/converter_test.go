package converter

import (
	"testing"
)

func TestNewImageConverter(t *testing.T) {
	conv := NewImageConverter()
	if conv == nil {
		t.Fatal("NewImageConverter returned nil")
	}
	if conv.Quality != 85 {
		t.Errorf("Expected default quality 85, got %d", conv.Quality)
	}
}

func TestSetQuality(t *testing.T) {
	conv := NewImageConverter()

	// Test valid quality
	conv.SetQuality(75)
	if conv.Quality != 75 {
		t.Errorf("Expected quality 75, got %d", conv.Quality)
	}

	// Test quality too low
	conv.SetQuality(-10)
	if conv.Quality != 1 {
		t.Errorf("Expected quality 1 (minimum), got %d", conv.Quality)
	}

	// Test quality too high
	conv.SetQuality(150)
	if conv.Quality != 100 {
		t.Errorf("Expected quality 100 (maximum), got %d", conv.Quality)
	}

	// Test boundary values
	conv.SetQuality(1)
	if conv.Quality != 1 {
		t.Errorf("Expected quality 1, got %d", conv.Quality)
	}

	conv.SetQuality(100)
	if conv.Quality != 100 {
		t.Errorf("Expected quality 100, got %d", conv.Quality)
	}
}
