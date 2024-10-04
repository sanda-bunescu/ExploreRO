package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/sanda-bunescu/ExploreRO/initializers"
)

type FirebaseService struct {
	client *auth.Client
}

func NewFirebaseService() (*FirebaseService, error) {
	client, err := initializers.FirebaseApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase Auth client")
	}
	return &FirebaseService{client: client}, err
}

func (f *FirebaseService) VerifyIDToken(ctx *gin.Context) (string, error) {
	idToken := ctx.Request.Header.Get("Authorization")
	if idToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No idToken provided"})
		fmt.Println("No idToken provided")
		return "", fmt.Errorf("no idToken provided")
	}
	if len(idToken) > 7 && idToken[:7] == "Bearer " {
		idToken = idToken[7:]
	}

	token, err := f.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify ID token."})
		return "", fmt.Errorf("failed to verify ID token")
	}

	log.Printf("Verified ID token: %v\n", token)
	return token.UID, nil
}

func (f *FirebaseService) GetUserByUID(ctx *gin.Context, firebaseUID string) *auth.UserRecord {
	userRecord, err := f.client.GetUser(context.Background(), firebaseUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Firebase user details."})
		fmt.Println("Failed to get Firebase user details.")
		return nil
	}
	return userRecord
}
