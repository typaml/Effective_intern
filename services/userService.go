package services

import (
	"Effective_intern/models"
	"Effective_intern/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: *repositories.NewUserRepository(),
	}
}

func (us *UserService) GetUsers() ([]models.User, error) {
	return us.userRepo.GetUsers()
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return us.userRepo.GetUserByID(id)
}

func (us *UserService) GetUserTasks(id uint) ([]models.Task, error) {
	return us.userRepo.GetUserTasks(id)
}

func (us *UserService) CreateUser(user *models.User) error {
	return us.userRepo.CreateUser(user)
}

func (us *UserService) UpdateUser(id uint, user *models.User) error {
	return us.userRepo.UpdateUser(id, user)
}

func (us *UserService) DeleteUser(id uint) error {
	return us.userRepo.DeleteUser(id)
}

func (us *UserService) StartTask(taskId uint) error {
	return us.userRepo.StartTask(taskId)
}
func (us *UserService) StopTask(taskId uint) error {
	return us.userRepo.StopTask(taskId)
}
