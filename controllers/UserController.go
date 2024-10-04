package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanda-bunescu/ExploreRO/initializers"
	"github.com/sanda-bunescu/ExploreRO/repositories"
	"github.com/sanda-bunescu/ExploreRO/services"
)

type UserController struct {
	FirebaseService *services.FirebaseService
	UserService     *services.UserService
}

func NewUserController() (*UserController, error) {
	// Initialize Firebase service
	firebaseService, err := services.NewFirebaseService()
	if err != nil {
		return nil, err
	}

	userRepo := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepo)

	return &UserController{
		FirebaseService: firebaseService,
		UserService:     userService,
	}, nil
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	firebaseUID, err := uc.FirebaseService.VerifyIDToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		fmt.Println("User unauthorized")
		return
	}
	//Get userRecord from firebase
	userRecord := uc.FirebaseService.GetUserByUID(ctx, firebaseUID)

	//Get user from the DB
	user, err := uc.UserService.FindUserByEmail(ctx, userRecord.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {

		user, err = uc.UserService.AddUserInDB(ctx, firebaseUID, userRecord)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})
	} else {

		//user exists
		if user.DeletedAt != nil {
			//re-register user
			err := uc.UserService.UpdateDeletedUser(ctx, user, firebaseUID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User reactivated", "user": user})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	firebaseUID, err := uc.FirebaseService.VerifyIDToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		fmt.Println("User unauthorized")
		return
	}

	user, err := uc.UserService.GetUserByFirebaseId(ctx, firebaseUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No user found in the DB"})
		fmt.Println("No user found in the DB")
		return
	} else {
		//user exists
		err := uc.UserService.SoftDelete(ctx, user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	}
}
