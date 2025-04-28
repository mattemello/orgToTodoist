package parse

import (
	"math"
	"strings"
)

func ParseAdd(strCont string) []string {
	indexMod := strings.Index(strCont, "modified: ")
	if indexMod == -1 {
		indexMod = math.MaxInt
	}
	indexAdd := strings.Index(strCont, "new file: ")
	if indexAdd == -1 {
		indexAdd = math.MaxInt
	}

	var whereToTake int
	whereToTake = min(indexMod, indexAdd)

	if whereToTake != math.MaxInt {
		usefulCont := strCont[whereToTake:]

		usefulCont = strings.ReplaceAll(usefulCont, "\n", "")
		usefulCont = strings.ReplaceAll(usefulCont, " ", "")
		usefulCont = strings.ReplaceAll(usefulCont, "\t", ":")

		separeCont := strings.Split(usefulCont, ":")

		var file []string

		for i := 0; i < len(separeCont); i += 2 {
			file = append(file, separeCont[i+1])
		}

		return file
	}

	return nil

}
