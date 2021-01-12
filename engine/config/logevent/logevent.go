package logevent

import (
	"os"
	"regexp"
	"strings"
	"time"
)

type LogEvent struct {
	Timestamp time.Time              `json:"timestamp"`
	Message   string                 `json:"message"`
	Tags      []string               `json:"tags"`
	Extra     map[string]interface{} `json:"-"`
}

type Config struct {
	SortMapKeys bool     `yaml:"sort_map_keys"`
	RemoveField []string `yaml:"remove_field"`

	jsonMarshal       func(v interface{}) ([]byte, error)
	jsonMarshalIndent func(v interface{}) ([]byte, error)
}

const TagsField = "tags"

const timeFormat = `2006-01-02T15:04:05.999999999Z`

var config *Config

func SetConfig(c *Config) {
	config = c
}

func init() {
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func (t *LogEvent) AddTag(tags ...string) {
	for _, tag := range tags {
		ftag := t.Format(tag)
		t.Tags = appendIfMissing(t.Tags, ftag)
	}
}

var (
	reCurrentTime = regexp.MustCompile(`%{\+([^}]+)}`)
	reEventTime   = regexp.MustCompile(`%{\+@([^}]+)}`)
	revar         = regexp.MustCompile(`%{([\w@\.]+)}`)
)

// FormatWithEnv format string with environment value, ex: %{HOSTNAME}
func FormatWithEnv(text string) (result string) {
	result = text

	matches := revar.FindAllStringSubmatch(result, -1)
	for _, submatches := range matches {
		field := submatches[1]
		value := os.Getenv(field)
		if value != "" {
			result = strings.Replace(result, submatches[0], value, -1)
		} else if field == "HOSTNAME" {
			if value, _ := os.Hostname(); value != "" {
				result = strings.Replace(result, submatches[0], value, -1)
			}
		}
	}

	return
}

// FormatWithCurrentTime format string with current time, ex: %{+2006-01-02}
func FormatWithCurrentTime(text string) (result string) {
	result = text

	matches := reCurrentTime.FindAllStringSubmatch(result, -1)
	for _, submatches := range matches {
		value := time.Now().Format(submatches[1])
		result = strings.Replace(result, submatches[0], value, -1)
	}

	return
}

// FormatWithEventTime format string with event time, ex: %{+@2006-01-02}
func FormatWithEventTime(text string, evevtTime time.Time) (result string) {
	result = text

	matches := reEventTime.FindAllStringSubmatch(result, -1)
	for _, submatches := range matches {
		value := evevtTime.Format(submatches[1])
		result = strings.Replace(result, submatches[0], value, -1)
	}

	return
}

// Format return string with current time / LogEvent field / ENV, ex: %{hostname}
func (t LogEvent) Format(format string) (out string) {
	out = format

	out = FormatWithEventTime(out, t.Timestamp)
	out = FormatWithCurrentTime(out)

	matches := revar.FindAllStringSubmatch(out, -1)
	for _, submatches := range matches {
		field := submatches[1]
		value := field // t.GetString(field)
		if value != "" {
			out = strings.Replace(out, submatches[0], value, -1)
		}
	}
	out = FormatWithEnv(out)
	return
}
