package lesson9

import (
	"fmt"
	"regexp"
)

func slugify(src string) string {
	//toLower := func(src string) string {
	//	return strings.ToLower(src)
	//}

	return ""
}

func Test() {
	re := regexp.MustCompile(`[a-z0-9-]`)
	test := "kwemf-*fwef"
	fmt.Println(re.MatchString(test))
}
