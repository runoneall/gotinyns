package records

import (
	"bytes"
	"strings"
)

func Find(name []byte, qtype uint16) []Record {
	key := bytes.ToLower(trimTrailingDot(name))

	mu.RLock()
	defer mu.RUnlock()

	if isReverseName(key) {
		return findReverse(key, qtype)
	}

	return findForward(key, qtype)
}

func findForward(key []byte, qtype uint16) []Record {
	var out []Record
	for _, r := range records {
		if !bytes.Equal(bytes.ToLower(trimTrailingDot(r.Name)), key) {
			continue
		}

		if qtype == 255 /* ANY */ || r.Type == qtype {
			out = append(out, r)
		}
	}

	return out
}

func findReverse(key []byte, qtype uint16) []Record {
	if qtype != 12 /* PTR */ && qtype != 255 {
		return nil
	}

	ipStr := reverseNameToIP(key)

	var out []Record
	for _, r := range records {
		if r.Type != 12 /* PTR */ {
			continue
		}

		rName := bytes.ToLower(trimTrailingDot(r.Name))
		if bytes.Equal(rName, key) {
			out = append(out, r)
			continue
		}

		if ipStr != "" && bytes.Equal(rName, []byte(ipStr)) {
			out = append(out, r)
		}
	}

	return out
}

func trimTrailingDot(name []byte) []byte {
	if len(name) > 0 && name[len(name)-1] == '.' {
		return name[:len(name)-1]
	}

	return name
}

func isReverseName(name []byte) bool {
	suffix := []byte(".in-addr.arpa")
	if !bytes.HasSuffix(name, suffix) {
		return false
	}

	return len(name) > len(suffix)
}

func reverseNameToIP(name []byte) string {
	s := string(name)
	const suffix = ".in-addr.arpa"
	if !strings.HasSuffix(s, suffix) {
		return ""
	}

	labels := strings.Split(strings.TrimSuffix(s, suffix), ".")
	if len(labels) != 4 {
		return ""
	}

	for i, j := 0, len(labels)-1; i < j; i, j = i+1, j-1 {
		labels[i], labels[j] = labels[j], labels[i]
	}

	for _, l := range labels {
		if !isNumeric(l) || len(l) == 0 || len(l) > 3 {
			return ""
		}
	}

	return strings.Join(labels, ".")
}

func isNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}
