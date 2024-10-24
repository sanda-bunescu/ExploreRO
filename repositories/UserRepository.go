package repositories

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanda-bunescu/ExploreRO/initializers"
	"github.com/sanda-bunescu/ExploreRO/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) GetUserByEmail(ctx *gin.Context, email string) (*models.Users, error) {
	var user models.Users
	userResult := ur.DB.Where("email = ?", email).First(&user)
	if userResult.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if userResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user from database."})
		return nil, fmt.Errorf("failed to query user from database")
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByFirebaseId(ctx *gin.Context, firebaseUID string) (*models.Users, error) {
	var user models.Users
	userResult := initializers.DB.Where("firebase_id = ?", firebaseUID).First(&user)
	if userResult.Error != nil && userResult.Error != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user from database."})
		return nil, fmt.Errorf("failed to query user from database")
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(ctx *gin.Context, user *models.Users) error {
	result := ur.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return fmt.Errorf("failed to create user")
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
	return nil
}

func (ur *UserRepository) UpdateDeletedUser(ctx *gin.Context, user *models.Users, firebaseUID string) error {
	user.DeletedAt = nil
	user.FirebaseId = firebaseUID //update to new uid

	result := initializers.DB.Save(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to re-reauthenticate user"})
		return fmt.Errorf("failed to re-reauthenticate user")
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User reactivated", "user": user})
	return nil
}

func (ur *UserRepository) SoftDelete(ctx *gin.Context, user *models.Users) error {
	currentTime := time.Now()
	user.DeletedAt = &currentTime

	result := initializers.DB.Save(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return fmt.Errorf("failed to delete user")
	}
	return nil
}
