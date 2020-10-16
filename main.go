package main

import (
	"fmt"
	"regexp"
)

const defaultName = "MAG5.1-160-236(740.S.48R.AS)"
const industryName = "INDUSTRY.3-160-236(740.G.BR2.AS)"
const ballName = "BALL400-060(840.N)"
const ledparkName = "LedPark02-028"

//var magRgx = regexp.MustCompile(`(([A-Z]{3,}\.?\w?[\d\.]+)-(\d+)-(\d)(\d+))\w?\((\d)(\d+)\.([A-Z])\.(\d*[A-Z]+\d*)\.([A-Z]+)\)`)

var magRgx = regexp.MustCompile(`(MAG|INDUSTRY|BALL|LedPark)\.?[\d\.]*-\d+(-\d{3})?(\w?\(\d{3}(\.[A-Z])?(\.\d*[A-Z]+\d*)?(\.[A-Z]+)?\))?`)

func main() {
	if magRgx.MatchString(defaultName) {
		// submatches := magRgx.FindAllStringSubmatch(defaultName, -1)[0][1:]
		// fmt.Println(submatches)
		fmt.Println("MAG completed")
	}

	if magRgx.MatchString(industryName) {
		// submatches := magRgx.FindAllStringSubmatch(industryName, -1)[0][1:]
		// fmt.Println(submatches)
		fmt.Println("MAG completed")
	}

	if magRgx.MatchString(ballName) {
		fmt.Println("BALL completed")
	}

	if magRgx.MatchString(ledparkName) {
		fmt.Println("LedPark completed")
	}
}
