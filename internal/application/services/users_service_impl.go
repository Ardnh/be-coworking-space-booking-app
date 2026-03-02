package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
)

type UserServiceImpl struct {
	repo repositories.UsersRepository
	log  *logrus.Logger
}

func NewUserService(repo repositories.UsersRepository, log *logrus.Logger) services.UserService {
	return &UserServiceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *UserServiceImpl) GetUserById(ctx context.Context, userId uuid.UUID) (*dto.UserDto, error) {

	user, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	userDto := mapper.UserEntityToDto(user)
	return userDto, err
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

	// Conver to entities
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userPassword := string(hashedPassword)
	user.Password = userPassword
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
func (s *UserServiceImpl) DeleteUser(ctx context.Context, userID uuid.UUID) error {

	return s.repo.DeleteUser(ctx, userID)
}
