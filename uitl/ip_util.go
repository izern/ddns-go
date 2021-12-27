package uitl

func init() {

}

func IsIP(str string) bool {

	if IsIPV4(str) || isIPV6(str) {
		return true
	}
	return false

}

func IsIPV4(str string) bool {
	return true
}

func isIPV6(str string) bool {
	return true
}
