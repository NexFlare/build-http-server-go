package response

type Status string

const (
	StatusOk Status = "200 OK"
	StatusCreated = "201 Created"
	StatusBadRequest = "400 Bad Request"
	StatusNotFound = "404 Not Found"
	StatusInternalServerError = "500 Internal Server Error"

)