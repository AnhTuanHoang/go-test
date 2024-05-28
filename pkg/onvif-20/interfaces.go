package onvif20

type Function interface {
	Request() interface{}
	Response() interface{}
}
