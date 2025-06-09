package serializer

import "singo/model"

// User User serializer
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser Serialize user
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse Serialize user response
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
