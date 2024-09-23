package data

type APIErrors struct {
	Errors []*APIError `json:"errors"`
}

func (errors *APIErrors) Status() int {
	if len(errors.Errors) > 0 {
		return errors.Errors[0].Status
	}
	return -1
}

type APIError struct {
	Status       int    `json:"status"`
	Title        string `json:"title"`
	ErrorDetails string `json:"details"`
}

type APIErrorMeta struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}

func (out *APIErrors) NewAPIError(status int, title string, err string) {
	out.Errors = append(out.Errors, &APIError{
		Status:       status,
		Title:        title,
		ErrorDetails: err,
	})
}
