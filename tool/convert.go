package tool

import "strconv"

func ToString(s interface{}) string {
	switch vlu := s.(type) {
	case int:
		return strconv.Itoa(vlu)
	case int64:
		return strconv.FormatInt(vlu, 10)
	case string:
		return vlu
	}
	return ""
}
