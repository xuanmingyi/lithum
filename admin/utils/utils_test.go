package utils

import (
	"fmt"
	"testing"
)

func TestLoadComment(t *testing.T) {
	result, err := LoadComment("name: test\ndisplay: 特色\n")
	fmt.Println(result["name"])
	fmt.Println(result["display"])
	fmt.Println(err)
}
