package mapping

import (
	"naro-app-be/internal/entity"
	"naro-app-be/internal/model"
)

func UserMappingRequestDto(user *model.CreateUserRequestDto) *entity.User {
	return &entity.User{
		Username: user.Username,
		Password: user.Password,
		Name:     user.Name,
	}
}

func UserMappingResponseDto(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}
}

func UsersMappingResponseDTO(user []entity.User) *[]model.UserResponse {

	var result []model.UserResponse

	for _, item := range user {
		result = append(result, model.UserResponse{
			Username:  item.Username,
			Password:  item.Password,
			ID:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return &result
}
