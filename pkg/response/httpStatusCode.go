package response

const (
	ErrCodeSuccess = 2001  // Success
	ErrCodeParam   = 20003 // Email invalid
)

// message
var msg = map[int]string{
	ErrCodeSuccess: "Success",
	ErrCodeParam:   "Email is invalid",
}
