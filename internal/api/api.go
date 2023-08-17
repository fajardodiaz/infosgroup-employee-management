package api

import (
	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *database.Store
	router *gin.Engine
}

func NewServer(store *database.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// routes
	v1 := router.Group("/api/v1")

	states := v1.Group("/states")
	{
		states.POST("/", server.createState)
		states.GET("/:id", server.GetStateById)
		// states.GET("/")
		// states.DELETE("/:id")
		// states.PUT("/:id")
	}

	// projects := v1.Group("/projects")
	// {
	// 	projects.GET("/")
	// 	projects.GET("/:id")
	// 	projects.POST("/")
	// 	projects.DELETE("/:id")
	// 	projects.PUT("/:id")
	// }

	// positions := v1.Group("/positions")
	// {
	// 	positions.GET("/")
	// 	positions.GET("/:id")
	// 	positions.POST("/")
	// 	positions.DELETE("/:id")
	// 	positions.PUT("/:id")
	// }

	// teams := v1.Group("/teams")
	// {
	// 	teams.GET("/")
	// 	teams.GET("/:id")
	// 	teams.POST("/")
	// 	teams.DELETE("/:id")
	// 	teams.PUT("/:id")
	// }

	// employees := v1.Group("/employees")
	// {
	// 	employees.GET("/")
	// 	employees.GET("/:id")
	// 	employees.POST("/")
	// 	employees.DELETE("/:id")
	// 	employees.PUT("/:id")
	// }

	// employees_project := v1.Group("/assignments")
	// {
	// 	// Select all the employees in a project
	// 	employees_project.GET("/project/:id/employees")
	// 	// Select all the projects assigned to a single employee
	// 	employees_project.GET("/employee/:id/projects")
	// 	// Assign an employee to a project
	// 	employees_project.POST("/employee/:id/project/:id")
	// 	// Delete an employee from a project
	// 	employees_project.DELETE("/employee/:id/project/:id")
	// }

	server.router = router
	return server
}

// start the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
