package routes

import (
	"stury-supplements/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.POST("/login", controllers.Login)

	// Todas as rotas dentro deste grupo exigirão autenticação
	protectedGroup := r.Group("/protected")
	protectedGroup.Use(controllers.AuthMiddleware())
	{
		protectedGroup.POST("/endpoint1", controllers.ProtectedHandler)
		// adição de outras rotas protegidas
	}

	// Todas as rotas abaixo exigirão autenticação
	r.Use(controllers.AuthMiddleware())
	{
		r.POST("/new", controllers.RegistraSuplementos)
	}

	// Rota de login (sem autenticação)

	r.Run(":8001")
}
