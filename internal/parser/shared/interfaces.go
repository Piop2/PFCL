package shared

type OnCompleteCallback func(result any)

type State interface {
	SetContext(ctx *Context)
	SetOnComplete(f OnCompleteCallback)
	Process(
		token rune,
	) (next State, isProcessed bool, err ErrPFCL)
	IsParsing() bool
}
