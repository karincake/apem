package apem

type extCall func()

var extraCalls []extCall

func (a *app) initExtCall() {
	for _, init := range extraCalls {
		init()
	}
}

func RegisterExtCall(e extCall) {
	extraCalls = append(extraCalls, e)
}
