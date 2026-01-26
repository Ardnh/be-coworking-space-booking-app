package dto

// ================ DTO ====================
type UserDto struct {
	UserID       string  `json:"user_id" validate:"required,uuid"`
	Email        string  `json:"email" validate:"required,email"`
	Username     string  `json:"username" validate:"required,min=3,max=50"`
	FullName     string  `json:"full_name" validate:"required"`
	PhoneNumber  string  `json:"phone_number" validate:"required"`
	UserType     string  `json:"user_type" validate:"required,oneof=admin customer staff"`
	ProfileImage string  `json:"profile_image"`
	Status       string  `json:"status" validate:"required,oneof=active inactive suspended banned"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at,omitempty"`
}

// ================ REQUEST DTO ====================
type CreateUserDto struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=5"`
	Username     string `json:"username" validate:"required,min=1,max=50"`
	FullName     string `json:"full_name" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"omitempty,url"`
}

type UpdateUserDto struct {
	Username     string `json:"username" validate:"required,min=1,max=50"`
	FullName     string `json:"full_name" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"omitempty,url"`
}

// ================ Params ====================
type GetUserParams struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"fullname"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}
