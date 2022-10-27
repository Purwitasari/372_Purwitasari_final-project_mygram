package routers

import (
	"MyGram/controllers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	authRouter := r.Group("/users")
	{
		// Create/Register User
		authRouter.POST("/register", controllers.UserRegister)

		// Login User
		authRouter.POST("/login", controllers.UserLogin)
	}

	userRouter := r.Group("/users")
	{

		// Get User For Profile
		userRouter.GET("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.UserGet)

		// Update Profile User
		userRouter.PUT("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.UserUpdate)

		// Delete Account user
		userRouter.DELETE("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		// Create New Photo
		photoRouter.POST("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.PhotoCreate)

		// Get Account Photo
		photoRouter.GET("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.PhotoGet)

		// Update Account photo
		photoRouter.PUT("/:photoId", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.PhotoUpdate)

		// Delete Photo Account
		photoRouter.DELETE("/:photoId", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.PhotoDelete)
	}

	commentRouter := r.Group("/comments")
	{
		// Create New Comment
		commentRouter.POST("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.CreateComment)

		// Get Comment
		commentRouter.GET("/:photoId", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.GetComment)

		// Update Comment
		commentRouter.PUT("/:commentId", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.UpdateComment)

		// Delete Comment
		commentRouter.DELETE("/:commentId", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.DeleteComment)
	}

	sosmedRouter := r.Group("/socialmedias")
	{
		// Create New Sosmed
		sosmedRouter.POST("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.CreateSosmed)

		// Get Account Sosmed
		sosmedRouter.GET("/", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.GetSosmed)

		// Update Account sosmed
		sosmedRouter.PUT("/:socialMediaId", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.UpdateSosmed)

		// Delete Account Sosmed
		sosmedRouter.DELETE("/:socialMediaId", middlewares.Authentication(), middlewares.UserAuthentication(), controllers.DeleteSosmed)
	}

	return r
}
