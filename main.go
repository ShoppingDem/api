package main

import (
	"os"

	_ "github.com/ShoppingDem/api/docs" // This line is important
	"github.com/ShoppingDem/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/okta/okta-jwt-verifier-golang"
)

// @title Shopping API
// @version 1.0
// @description This is a sample shopping API server with Okta authentication.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := routes.SetupRouter()

	// Set up Okta middleware
	r.Use(authMiddleware())

	r.Run(":8080")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}

		jwtVerifierSetup := jwtverifier.JwtVerifier{
			Issuer:           os.Getenv("OKTA_ISSUER"),
			ClaimsToValidate: map[string]string{"aud": "api://default", "cid": os.Getenv("OKTA_CLIENT_ID")},
		}

		verifier := jwtVerifierSetup.New()

		claims, err := verifier.VerifyAccessToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
