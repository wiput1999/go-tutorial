package slice

func Slice(a []string, b []string) []string {

	var s = []string{}

	for i := 0; i < len(a); i++ {
		s = append(s, a[i], b[i])
	}

	return s
}
