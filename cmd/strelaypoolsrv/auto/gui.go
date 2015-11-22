package auto

import (
	"encoding/base64"
)

const (
	AssetsBuildDate = "Sun, 22 Nov 2015 22:45:35 GMT"
)

func Assets() map[string][]byte {
	var assets = make(map[string][]byte, 1)

	assets["index.html"], _ = base64.StdEncoding.DecodeString("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")
	return assets
}
