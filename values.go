package mws

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

//NewValues NewValue
func NewValues(params ...map[string]string) Values {
	return Values{}
}

//ActionValues NewValue
func ActionValues(action string) Values {
	return Values{keyAction: action}
}

//Values Values
type Values map[string]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	return v[key]
}

//Set set string value
func (v Values) Set(name, value string) {
	if v != nil && name != "" && value != "" {
		v[name] = value
	}
}

// Delete deletes the values associated with name.
func (v Values) Delete(name string) {
	if v != nil {
		delete(v, name)
	}
}

// Deletes deletes the values associated with name.
func (v Values) Deletes(name string) {
	if v != nil {
		for key := range v {
			if strings.HasPrefix(key, name+".") {
				delete(v, key)
			}
		}
	}
}

//SetTime setTime
func (v Values) SetTime(name string, t time.Time) {
	if !t.IsZero() {
		v.Set(name, t.Format(time.RFC3339))
	}
}

//SetTimestamp SetTimestamp
func (v Values) SetTimestamp(name string, timestamp int64) {
	if timestamp > 0 {
		v.SetTime(name, time.Unix(timestamp, 0))
	}
}

//SetInt SetInt
func (v Values) SetInt(name string, value int64) {
	if value > 0 {
		v.Set(name, strconv.FormatInt(value, 10))
	}
}

//SetBool 设置bool
func (v Values) SetBool(name string, value int8) {
	if value > 0 {
		v.Set(name, strconv.FormatBool(value == 1))
	}
}

//Sets 设置集合
func (v Values) Sets(name string, values ...string) {
	if len(values) > 0 {
		for i, value := range values {
			v.Set(fmt.Sprintf("%s.%d", name, i+1), value)
		}
	}
}

//SetVersion set version value
func (v Values) SetVersion(version string) {
	if v != nil && version != "" {
		v[keyVersion] = version
	}
}

//SetAction set action value
func (v Values) SetAction(action string) {
	if v != nil && action != "" {
		v[keyAction] = action
	}
}

//SetAll 用新值覆盖
func (v Values) SetAll(params ...Values) {
	for _, param := range params {
		for name, value := range param {
			v.Set(name, value)
		}
	}
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (v Values) Encode() string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i > 0 {
			buf.WriteRune('&')
		}
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v[k]))
	}
	return buf.String()
}
