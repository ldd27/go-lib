package validation

import "strings"

// IsChinaMobile 验证手机号
func IsChinaMobile(mobile string) bool {
	if len(mobile) != 11 || !strings.HasPrefix(mobile, "1") {
		return false
	}
	return true
}
