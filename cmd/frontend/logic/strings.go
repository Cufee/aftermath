package logic

func StringIfElse(onTrue, onFalse string, condition bool) string {
	if condition {
		return onTrue
	}
	return onFalse
}
