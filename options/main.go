package options

type Option interface {
	GetType() string
	GetOptions() string
	ParseOptions(string)

	GetValue() string
	GetValueBytes() uint64
	GetValueInt() uint64
	GetOperation() string
	GetCopy() Option
}

func CompareStrWithOp(left string, right string, checkOp string) bool {
	if checkOp == "=" && left == right {
		return true
	}
	if checkOp == "==" && left == right {
		return true
	}
	if checkOp == "!=" && left != right {
		return true
	}
	return false
}

func CompareIntWithOp(left uint64, right uint64, checkOp string) bool {
	if checkOp == ">" && left > right {
		return true
	}
	if checkOp == ">=" && left >= right {
		return true
	}
	if checkOp == "<" && left < right {
		return true
	}
	if checkOp == "<=" && left <= right {
		return true
	}
	if checkOp == "==" && left == right {
		return true
	}
	if checkOp == "=" && left == right {
		return true
	}
	if checkOp == "!=" && left != right {
		return true
	}
	return false
}
