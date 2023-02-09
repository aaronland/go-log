package log

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestEmit(t *testing.T) {

	var buf bytes.Buffer
	buf_wr := bufio.NewWriter(&buf)

	logger := log.New(buf_wr, "", 0)
	emit(DEBUG_PREFIX, logger, "Hello, %s", "testing")

	buf_wr.Flush()

	expected := fmt.Sprintf("%s Hello, testing", DEBUG_PREFIX)
	log_str := strings.TrimSpace(buf.String())

	if log_str != expected {
		t.Fatalf("Unexpected log string. Expected '%s' but got '%s'", expected, log_str)
	}

	buf.Reset()

	err := fmt.Errorf("This is an error")

	emit(DEBUG_PREFIX, logger, err)

	buf_wr.Flush()

	expected = fmt.Sprintf("%s This is an error", DEBUG_PREFIX)
	log_str = strings.TrimSpace(buf.String())

	if log_str != expected {
		t.Fatalf("Unexpected log string. Expected '%s' but got '%s'", expected, log_str)
	}

	//

	buf.Reset()

	err2 := fmt.Errorf("This is a second error")

	r, w, p_err := os.Pipe()

	if p_err != nil {
		t.Fatalf("Failed to create pipe r/wr, %v", p_err)
	}

	stdout := os.Stdout
	stderr := os.Stderr

	os.Stdout = w
	os.Stderr = w

	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()

	Debug(err2)

	w.Close()

	_, cp_err := io.Copy(buf_wr, r)

	if cp_err != nil {
		t.Fatalf("Failed to copy pipe r, %v", cp_err)
	}

	buf_wr.Flush()

	expected = fmt.Sprintf("%s This is a second error", DEBUG_PREFIX)
	log_str = strings.TrimSpace(buf.String())

	if !strings.HasSuffix(log_str, expected) {
		t.Fatalf("Unexpected log string. Expected '%s' but got '%s'", expected, log_str)
	}

}

func TestDebug(t *testing.T) {

	var buf bytes.Buffer
	buf_wr := bufio.NewWriter(&buf)

	logger := log.New(buf_wr, "", 0)
	Debug(logger, "Hello, %s", "testing")

	buf_wr.Flush()

	expected := fmt.Sprintf("%s Hello, testing", DEBUG_PREFIX)
	log_str := strings.TrimSpace(buf.String())

	if log_str != expected {
		t.Fatalf("Unexpected log string. Expected '%s' but got '%s'", expected, log_str)
	}

}

func TestTestMinLevel(t *testing.T) {

	var buf bytes.Buffer
	buf_wr := bufio.NewWriter(&buf)
	logger := log.New(buf_wr, "", 0)

	//

	SetMinLevelWithPrefix(ERROR_PREFIX)

	Debug(logger, "Hello, %s", "testing")

	buf_wr.Flush()

	expected := ""
	log_str := strings.TrimSpace(buf.String())

	if log_str != expected {
		t.Fatalf("Unexpected log string. Expected '%s' but got '%s'", expected, log_str)
	}

	buf.Reset()

	//

	UnsetMinLevel()

	Debug(logger, "Hello, %s", "testing")

	buf_wr.Flush()

	expected = fmt.Sprintf("%s Hello, testing", DEBUG_PREFIX)
	log_str = strings.TrimSpace(buf.String())

	if log_str != expected {
		t.Fatalf("Unexpected log string. Expected '%s' but got '%s'", expected, log_str)
	}

	buf.Reset()

	//

	SetMinLevel(WARNING_LEVEL)

	Debug(logger, "Hello, %s", "testing")

	buf_wr.Flush()

	expected = ""
	log_str = strings.TrimSpace(buf.String())

	if log_str != expected {
		t.Fatalf("Unexpected log string. Expected '%s' but got '%s'", expected, log_str)
	}

}
