package error

type IllegalArg interface {
	error
}

type IllegalArgCondition struct {
	msg string
}

func (e IllegalArgCondition) Error() string {
	return e.msg
}

func doSth2() IllegalArg {
	return nil
}
