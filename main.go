package main

import (
	"fmt"
	"regexp"
)

const defaultName = "MAG5.1-160-236(740.S.48R.AS)"
const industryName = "INDUSTRY.3-160-236(740.G.BR.AS)"

var magRgx = regexp.MustCompile(`(([A-Z]{3,}\.?[\d\.]+)-(\d+)-(\d)(\d+))\w?\((\d)(\d+)\.([A-Z])\.(\d*[A-Z]+\d*)\.([A-Z]+)\)`)

func main() {
	if magRgx.MatchString(defaultName) {
		submatches := magRgx.FindAllStringSubmatch(defaultName, -1)[0][1:]
		fmt.Println(submatches)
	}

	if magRgx.MatchString(industryName) {
		submatches := magRgx.FindAllStringSubmatch(industryName, -1)[0][1:]
		fmt.Println(submatches)
	}
}
