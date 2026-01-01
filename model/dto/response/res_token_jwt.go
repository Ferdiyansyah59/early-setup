package response

type TokenResponseDTO struct {
	Token string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}