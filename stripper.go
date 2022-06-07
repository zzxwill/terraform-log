package stripper

import "regexp"

// StripColor removes color codes from Terraform log
func StripColor(log string) string {
	var re = regexp.MustCompile(`\x1b\[[0-9;]*m`)
	str := re.ReplaceAllString(log, "")
	return str
}
