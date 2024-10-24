package services

import (
	"fmt"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/sanda-bunescu/ExploreRO/models"
	"github.com/sanda-bunescu/ExploreRO/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (us *UserService) FindUserByEmail(ctx *gin.Context, email string) (*models.Users, error) {
	user, err := us.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by email: %v", err)
	}
	return user, nil
}

func (us *UserService) AddUserInDB(ctx *gin.Context, firebaseUID string, userRecord *auth.UserRecord) (*models.Users, error ){

	user := &models.Users{
		FirebaseId: firebaseUID,
		Name:       userRecord.DisplayName,
		Email:      userRecord.Email,
		CreatedAt:  time.Now(),
	}
	if err := us.UserRepo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return user, nil
}

func (us *UserService) UpdateDeletedUser(ctx *gin.Context, user *models.Users, firebaseUID string) error {
	if err := us.UserRepo.UpdateDeletedUser(ctx, user, firebaseUID); err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (us *UserService) GetUserByFirebaseId(ctx *gin.Context, firebaseUID string) (*models.Users, error){
	user, err := us.UserRepo.GetUserByFirebaseId(ctx, firebaseUID)
	if err != nil{
		return nil, fmt.Errorf("failed to get user")
	}
	return user, nil
}

func (us *UserService) SoftDelete(ctx *gin.Context, user *models.Users) error {
	err := us.UserRepo.SoftDelete(ctx, user)
	if err != nil{
		return fmt.Errorf("failed to get user")
	}
	return nil
}


