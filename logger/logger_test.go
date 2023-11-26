package logger

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"
)

type Source struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}
type LogTest struct {
	Time   time.Time `json:"time"`
	Level  string    `json:"level"`
	Source Source    `json:"source"`
	Msg    string    `json:"msg"`
	Args1  string    `json:"args1"`
	Args2  int       `json:"args2"`
}

func TestInitLogger(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "logger-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	cfg := Config{
		Filename:    tmpDir + "/test.json",
		MaxSize:     1,
		MaxBackups:  2,
		MaxAge:      1,
		Compress:    true,
		LocalTime:   true,
		LogLevel:    "INFO",
		AddSource:   true,
		HandlerType: "json",
	}

	InitLogger(cfg)

	Logger.Info("This is message", slog.String("args1", "one"), slog.Int("args2", 1))
	Logger.Debug("This is a debug message", slog.String("args1", "one"), slog.Int("args2", 1))
	Logger.Error("This is message", slog.String("args1", "one"), slog.Int("args2", 1))
	Logger.Warn("This is message", slog.String("args1", "one"), slog.Int("args2", 1))

	data, err := os.Open(cfg.Filename)
	if err != nil {
		t.Error(err)
	}

	scanner := bufio.NewScanner(data)

	var logs []LogTest

	for scanner.Scan() {

		line := scanner.Text()

		var log LogTest

		err := json.Unmarshal([]byte(line), &log)
		if err != nil {
			t.Error(err)
		}

		logs = append(logs, log)
	}

	for _, l := range logs {

		if !strings.Contains(l.Msg, "This is") {
			t.Error("Log file does not contain the info message")
		}

		if strings.Contains(l.Msg, "This is a debug message") {
			t.Error("Log file contains the debug message, which should be filtered out by the log level")
		}

	}

	// Check if the log file has the correct log level and source information

	for _, l := range logs {

		if l.Level != "INFO" && l.Level != "WARN" && l.Level != "ERROR" {
			t.Error("Log file has invalid log level")
		}
		if !strings.HasSuffix(l.Source.File, "logger_test.go") {
			t.Error("Log file does not have the correct source information")
		}

	}

	// Log some more messages to trigger the rotation and compression of the log files
	for i := 0; i < 10000; i++ {
		Logger.Info("This is message", slog.String("args1", "one"), slog.Int("args2", 1))
	}

	// Check if the log files are rotated and compressed as expected
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Error(err)
	}
	if len(files) != 3 {
		t.Error("Log directory does not have the expected number of files")
	}
	for _, file := range files {
		if file.Name() != "test.json" && !strings.HasSuffix(file.Name(), ".gz") {
			t.Error("Log file is not compressed")
		}
	}

	// Log a success message
	t.Log("TestInitLogger passed")
}
