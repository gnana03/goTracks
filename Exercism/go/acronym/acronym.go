package acronym

import
(
	"regexp"
	"fmt"
	"strings"
)
	
// This returns the abbreviated string
func Abbreviate(s string) string{
	re := regexp.MustCompile(`\s\w|\-\w|\_\w`)
	list:= re.FindAllString(s, -1)
	b := strings.Split(s,"")
	finalString := b[0]
	for i := 0;i < len(list);i++{
	  finalString = strings.ToUpper(finalString) + strings.ToUpper(list[i][1:])
	}
	fmt.Println(finalString)
	return finalString
}
