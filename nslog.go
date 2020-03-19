package nslog

type IosNSLogHandler interface {
	Log(str string)
}

var _logHandler IosNSLogHandler

func SetIosNSLogHandler(nslog IosNSLogHandler) {
	_logHandler = nslog
}

func NSLog(str string) {
	_logHandler.Log(str)

}
