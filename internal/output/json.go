package output

import (
	"encoding/json"
	"io"
)

func AsJson(i any, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(i)
}
