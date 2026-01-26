package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
)

type UserServiceImpl struct {
	repo repositories.UsersRepository
}

func NewUserService(repo repositories.UsersRepository) services.UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

// User Type Admin & Vendor
func (s *UserServiceImpl) GetAllUsers(ctx context.Context, userParams *dto.GetUserParams) ([]*dto.UserDto, int, error) {
	users, totalItems, err := s.repo.GetAllUsers(
		ctx,
		userParams.Username,
		userParams.Email,
		userParams.FullName,
		userParams.Limit,
		userParams.Offset,
		userParams.SortBy,
		userParams.SortOrder,
	)

	if err != nil {
		return nil, 0, err
	}

	usersListDto := mapper.UserListEntityToDto(users)
	return usersListDto, totalItems, nil
}

// User Type Customer
func (s *UserServiceImpl) CreateUser(ctx context.Context, user *dto.CreateUserDto) error {

	// convert to entities
	userEntity := mapper.CreateUserDtoToDomain(user)
	return s.repo.CreateUser(ctx, &userEntity)
}

// User Type Customer
func (s *UserServiceImpl) UpdateUser(ctx context.Context, userID uuid.UUID, user *dto.UpdateUserDto) error {

	// convert to entities
	userEntity := mapper.UpdateUserDtoToDomain(userID, user)
	return s.repo.UpdateUser(ctx, &userEntity)
}

// User Type Admin
func (s *UserServiceImpl) DeleteUser(ctx context.Context, userID string) error {

	return s.repo.DeleteUser(ctx, userID)
}
