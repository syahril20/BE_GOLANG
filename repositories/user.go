package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// Kontrak (PARAM DAN)
type UserRepository interface {
	FindUser() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	FindUserId(Id int) (models.User, error)
	DeleteUser(Id int, user models.User) (models.User, error)
	UpdateUser(Id int, user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repositories {
	return &repositories{db}
}
func (r *repositories) FindUser() ([]models.User, error) {
	var Users []models.User
	err := r.db.Find(&Users).Error
	// err := r.db.Raw("SELECT * FROM users").Scan(&user).Error

	return Users, err
}

func (r *repositories) FindUserId(Id int) (models.User, error) {
	var User models.User
	err := r.db.First(&User, Id).Error
	// err := r.db.Raw("SELECT * FROM users where id=?", Id).Scan(&User).Error

	return User, err
}
func (r *repositories) DeleteUser(Id int, user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error
	// err := r.db.Raw("DELETE FROM users WHERE id=?", Id).Scan(&user).Error

	return user, err
}

func (r *repositories) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	// err := r.db.Exec("INSERT INTO users(name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Error

	return user, err
}
func (r *repositories) UpdateUser(Id int, user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	// err := r.db.Exec("INSERT INTO users(name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Error

	return user, err
}
