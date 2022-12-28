package actioncyclicrotation

func DoTask(sl []int, shift int) []int {
	length := len(sl)
	rounds := shift / length
	if rounds > 0 {
		shift = shift - rounds*length
	}
	return append(sl[shift:], sl[:shift]...)
}
