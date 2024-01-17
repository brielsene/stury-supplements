package controllers

import (
	"fmt"
	"stury-supplements/database"
	"stury-supplements/models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RegistraSuplementos(c *gin.Context) {
	var suplementos models.Suplementos
	c.ShouldBindJSON(&suplementos)
	database.DB.Create(&suplementos)
	c.JSON(http.StatusCreated, &suplementos)
}

func Login(c *gin.Context) {
	fmt.Println("Entrei no login")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	fmt.Println("Você digitou: ", user)

	if err := authenticate(user.Usuario, user.Senha); err != nil {
		fmt.Println("Erro durante a autenticação:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	fmt.Println("Depois de chamar authenticate")

	tokenString, err := generateToken(user.Usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authenticate(Usuario, Senha string) error {
	fmt.Println("Entrando authenticate")
	var user models.User
	if err := database.DB.Where("Usuario = ?", Usuario).First(&user).Error; err != nil {
		fmt.Println("Erro ao buscar usuário no banco de dados:", err)
		return err
	}

	// Verifique se as senhas coincidem
	if user.Senha != Senha {
		fmt.Println("Senha incorreta para o usuário", Usuario)
		return fmt.Errorf("Senha incorreta")
	}

	fmt.Println("Usuário autenticado com sucesso:", user.Usuario)
	return nil
}

func generateToken(Usuario string) (string, error) {
	claims := TokenClaims{
		Usuario: Usuario,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Certifique-se de usar a mesma chave aqui que você usou para assinar o token
			return []byte("your-secret-key"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*TokenClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("Usuario", claims.Usuario)
		c.Next()
	}
}

func ProtectedHandler(c *gin.Context) {
	Usuario, exists := c.Get("Usuario")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Usuario from context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s!", Usuario)})
}

type TokenClaims struct {
	Usuario string `json:"usuario"`
	jwt.StandardClaims
}
