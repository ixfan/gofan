package web

import "strings"

var baseCode = 1
var baseCodeMultiple = 0

const (
	InternalServerError = iota + 1
	GatewayError
	TransactionError
	ArgError
	RecordNotFoundError
	BusinessError
)

type ServiceError struct {
	Code    int
	Message string
}

func (serviceError *ServiceError) Error() string {
	return serviceError.Message
}

func ThrowError(serviceErrorNo int, attachMessages ...string) error {
	switch serviceErrorNo {
	case InternalServerError:
		return &ServiceError{
			Code:    501 + baseCode*baseCodeMultiple,
			Message: strings.Join(append([]string{""}, attachMessages...), ""),
		}
	case GatewayError:
		return &ServiceError{
			Code:    502 + baseCode*baseCodeMultiple,
			Message: strings.Join(append([]string{""}, attachMessages...), ""),
		}
	case TransactionError:
		return &ServiceError{
			Code:    503 + baseCode*baseCodeMultiple,
			Message: strings.Join(append([]string{""}, attachMessages...), ""),
		}
	case ArgError:
		return &ServiceError{
			Code:    504 + baseCode*baseCodeMultiple,
			Message: strings.Join(append([]string{""}, attachMessages...), ""),
		}
	case RecordNotFoundError:
		return &ServiceError{
			Code:    505 + baseCode*baseCodeMultiple,
			Message: strings.Join(append([]string{""}, attachMessages...), ""),
		}
	case BusinessError:
		return &ServiceError{
			Code:    506 + baseCode*baseCodeMultiple,
			Message: strings.Join(append([]string{""}, attachMessages...), ""),
		}
	default:
		return &ServiceError{
			Code:    500 + baseCode*baseCodeMultiple,
			Message: strings.Join(append([]string{""}, attachMessages...), ""),
		}
	}
}

func NewError(code int, message string) error {
	return &ServiceError{
		Code:    code,
		Message: message,
	}
}
