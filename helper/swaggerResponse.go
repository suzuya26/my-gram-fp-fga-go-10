package helper

type UserRegisterResponse struct {
	Message  string `json:"message"`
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type CreateSosmedResponse struct {
	Message        string `json:"message"`
	Id             uint   `json:"id"`
	Name           string `json:"sosmed"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}

type UpdateSosmedResponse struct {
	Message        string `json:"message"`
	Id             uint   `json:"id"`
	Name           string `json:"sosmed"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}

type ErrorResponse struct {
	Err     string `json:"err"`
	Message string `json:"message"`
}
