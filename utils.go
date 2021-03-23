package mws

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
)

// md5->base64
func FeedContentMd5(body []byte) string {
	h := md5.New()
	h.Write(body)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func JsonMarshalIndentToString(v interface{}) string {
	bf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(bf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	encoder.Encode(v)
	return bf.String()
}
