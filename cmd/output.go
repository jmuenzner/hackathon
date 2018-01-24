package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// dumpJSON outputs json to stdout, optionally pretty printing it
func dumpJSON(value interface{}, human bool) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if human {
		var out bytes.Buffer
		if err = json.Indent(&out, b, " ", "\t"); err != nil {
			return err
		}

		fmt.Println(out.String())
		return nil
	}

	fmt.Println(string(b))
	return nil
}
