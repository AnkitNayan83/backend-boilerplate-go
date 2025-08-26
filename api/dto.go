package api

// Generic API response structure
type ApiResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// Helper function to create success response
func NewSuccessResponse(data any) ApiResponse {
	return ApiResponse{
		Success: true,
		Data:    data,
	}
}

// Helper function to create error response
func NewErrorResponse(message string) ApiResponse {
	return ApiResponse{
		Success: false,
		Message: message,
	}
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}
