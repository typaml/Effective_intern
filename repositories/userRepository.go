package repositories

import (
	"Effective_intern/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Task{})

	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := ur.db.First(&user, id).Error
	return &user, err
}

func (ur *UserRepository) GetUserTasks(id uint) ([]models.Task, error) {
	var tasks []models.Task
	err := ur.db.Where("user_id = ?", id).Find(&tasks).Error
	return tasks, err
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) UpdateUser(id uint, user *models.User) error {
	return ur.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (ur *UserRepository) DeleteUser(id uint) error {
	return ur.db.Delete(&models.User{}, id).Error
}

func (ur *UserRepository) StartTask(taskID uint) error {
	return ur.db.Model(&models.Task{}).Where("id = ?", taskID).Updates(map[string]interface{}{"start_time": time.Now()}).Error
}
func (ur *UserRepository) StopTask(taskID uint) error {
	return ur.db.Model(&models.Task{}).Where("id = ?", taskID).Updates(map[string]interface{}{"end_time": time.Now()}).Error
}
