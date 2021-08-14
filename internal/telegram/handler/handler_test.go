package handler

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGen(t *testing.T) {
	var msg = "this is a test with a lot of words. Will this work? I hope so. 1234123 $asdf ⁴23:test"

	re := regexp.MustCompile(`\s|[\W]`)
	x := re.Split(msg, -1)

	for i := 0; i < len(x); i++ {
		var phrase = make([]string, 12)
		for j := 0; j < 12; j++ {
			if i+j < len(x) {
				phrase[j] = x[j+i]
			} else {
				return
			}
		}
		fmt.Println(phrase)
	}

}
