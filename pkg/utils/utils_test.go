package utils

import (
	"bytes"
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	// Create a new Logger with Debug set to true.
	logger := Logger{Debug: true}

	// Create a buffer to capture log output.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Test Debugf.
	logger.Debugf("Debug message: %s", "unit")
	if !bytes.Contains(buf.Bytes(), []byte("Debug message: unit")) {
		t.Errorf("expected debug message, got %q", buf.String())
	}

	// Test Infof.
	logger.Infof("Info message: %s", "test")
	if !bytes.Contains(buf.Bytes(), []byte("Info message: test")) {
		t.Errorf("expected info message, got %q", buf.String())
	}

	// Test Warnf.
	logger.Warnf("Warning message: %s", "hi")
	if !bytes.Contains(buf.Bytes(), []byte("Warning message: hi")) {
		t.Errorf("expected warning message, got %q", buf.String())
	}

	// Test Errorf.
	logger.Errorf("Error message: %s", "you")
	if !bytes.Contains(buf.Bytes(), []byte("Error message: you")) {
		t.Errorf("expected error message, got %q", buf.String())
	}

}
