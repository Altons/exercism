package strand

import "strings"

//ToRNA converts
func ToRNA(dna string) string {
	var sb strings.Builder
	for _, str := range dna {
		char := string(str)
		switch {
		case char == "G":
			sb.WriteString("C")
		case char == "C":
			sb.WriteString("G")
		case char == "T":
			sb.WriteString("A")
		case char == "A":
			sb.WriteString("U")

		}
	}
	return sb.String()

}
