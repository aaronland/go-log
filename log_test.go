package log

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

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
