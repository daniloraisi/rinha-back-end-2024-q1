package errors

import "fmt"

type BadRequest struct {
	Reason string
}

func (n *BadRequest) Error() string {
	if n.Reason != "" {
		return "Bad Request: " + n.Reason
	}

	return "Bad Request"
}

type NotFound struct {
	Resource string
}

func (n *NotFound) Error() string {
	return fmt.Sprintf("%s not found", n.Resource)
}
