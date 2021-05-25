package registry

const (
	AppNotFoundCode =-1
)

var (
	AppNotFoundError = NewRegistryError(AppNotFoundCode,"app not found")
)

type RegistryError struct {

	Code int32
	Message string
}

// new a registry error
func NewRegistryError(code int32,message string)*RegistryError  {
	return &RegistryError{
		Code:    code,
		Message: message,
	}
}

func (e RegistryError)Error()string  {
	return e.Message
}