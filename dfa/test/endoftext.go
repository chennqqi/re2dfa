// Code generated by re2dfa (https://github.com/opennota/re2dfa).

package test

import "unicode/utf8"

//func isWordChar(r byte) bool {
//        return 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || '0' <= r && r <= '9' || r == '_'
//}

func matchEndOfText(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	_, _, _ = r, rlen, i
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 97:
		goto s2
	}
	return
s2:
	switch {
	case i == len(s):
		end = i
		return
	}
	return
}
