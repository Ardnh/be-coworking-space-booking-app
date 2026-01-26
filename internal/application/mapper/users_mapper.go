package mapper

import (
	"time"

	"github.com/google/uuid"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/pkg/constants"
)

func UserListEntityToDto(entities []*entities.Users) []*dto.UserDto {
	var dtos []*dto.UserDto
	for _, entity := range entities {
		dtos = append(dtos, UserEntityToDto(entity))
	}
	return dtos
}

func UserEntityToDto(entity *entities.Users) *dto.UserDto {

	return &dto.UserDto{
		UserID:       entity.UserID.String(),
		Email:        entity.Email,
		Username:     entity.Username,
		FullName:     entity.FullName,
		PhoneNumber:  entity.PhoneNumber,
		ProfileImage: entity.ProfileImage,
		Status:       entity.Status,
		UserType:     entity.UserType,
		CreatedAt:    entity.CreatedAt.String(),
		UpdatedAt:    entity.UpdatedAt.String(),
	}
}

func CreateUserDtoToDomain(dto *dto.CreateUserDto) entities.Users {

	now := time.Now()
	return entities.Users{
		Email:        dto.Email,
		Password:     dto.Password,
		Username:     dto.Username,
		FullName:     dto.FullName,
		PhoneNumber:  dto.PhoneNumber,
		ProfileImage: dto.ProfileImage,
		Status:       constants.UserStatusActive,
		UserType:     constants.UserTypeCustomer,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func UpdateUserDtoToDomain(userID uuid.UUID, dto *dto.UpdateUserDto) entities.Users {

	now := time.Now()
	return entities.Users{
		UserID:       userID,
		Username:     dto.Username,
		FullName:     dto.FullName,
		PhoneNumber:  dto.PhoneNumber,
		ProfileImage: dto.ProfileImage,
		UpdatedAt:    now,
	}
}
