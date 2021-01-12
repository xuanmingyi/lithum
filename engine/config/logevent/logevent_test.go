package logevent

import (
	"fmt"
	"testing"
	"time"
)

var now time.Time

func setup() {
	now = time.Now()
}

func TestFormatWithCurrentTime(t *testing.T) {
	fmt.Println(now)
	resutl := FormatWithCurrentTime("%{+2006-01-02 15:04:05}")
	fmt.Println(resutl)
}

func TestFormatWithEventTime(t *testing.T) {
	result := FormatWithEventTime("%{+@2006-01-02}", time.Now())
	fmt.Println(result)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}
