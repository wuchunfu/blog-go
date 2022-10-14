package util

import (
	"github.com/dlclark/regexp2"
	"strings"
)

var HTMLUtil htmlUtil

type htmlUtil struct {
}

func (htmlUtil) Filter(src string) string {
	re := regexp2.MustCompile(`(?!<(img|p|span|h1|h2|h3|h4|h5|h6).*?>)<.*?>`, 0)
	src, _ = re.Replace(src, "", 0, len(src))
	return strings.TrimSpace(src)
}
