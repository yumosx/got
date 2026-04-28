package lv_logic

/**
 * 自定义三元表达式
 */
func IfTrue(condition bool, trueResult, falseResult any) any {
	if condition {
		return trueResult
	}
	return falseResult
}
