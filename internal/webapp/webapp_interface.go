package webapp

type WebApp interface {
	Start() error
	Stop() error
}
