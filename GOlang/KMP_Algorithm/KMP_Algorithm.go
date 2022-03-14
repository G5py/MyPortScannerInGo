package algorithmStudy

func KmpSearch(str string, target string) (bool, int) {
	if len(target) > len(str) {
		return false, 0
	}

	pi := newPi(target)

	iter := 0
	for iter <= len(str)-len(target) {

		for str[iter] != target[0] && iter <= len(str)-len(target) {
			iter++
		}

		matchCount := 1
		for matchCount < len(target) {
			if target[matchCount] != str[iter+matchCount] {
				break
			}
			matchCount++
		}
		if matchCount == len(target) {
			return true, iter
		}

		iter += matchCount - pi[matchCount]
	}

	return false, 0
}

func newPi(target string) []int {
	pi := make([]int, len(target)+1)
	pi[0] = 0
	pi[1] = 0

	for i := 2; i <= len(target); i++ { // each process fills a cell of the slice
		pi[i] = func(str string) int {
			length := len(str)
			count := 1
			for count <= length/2 {
				if str[:count] != str[length-count:] {
					break
				}
				count++
			}

			return count - 1
		}(target[:i])
	}

	return pi
}
