package ami

// Response maps a string key to a list of values.
type Response map[string][]string

// Get gets the first value associated with the given key.
func (r Response) Get(key string) string {
	if r == nil {
		return ""
	}
	rs := r[key]
	if len(rs) == 0 {
		return ""
	}
	return rs[0]
}
