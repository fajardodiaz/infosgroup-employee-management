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
		states.GET("/", server.getStates)
		states.GET("/:id", server.getStateById)
		states.DELETE("/:id", server.deleteState)
		states.PUT("/:id", server.updateState)
	}

	projects := v1.Group("/projects")
	{
		projects.GET("/", server.getProjects)
		projects.GET("/:id", server.getProjectById)
		projects.POST("/", server.createProject)
		projects.DELETE("/:id", server.deleteProject)
		projects.PUT("/:id", server.updateProject)
	}

	positions := v1.Group("/positions")
	{
		positions.GET("/", server.getPositions)
		positions.GET("/:id", server.getPositionById)
		positions.POST("/", server.createPosition)
		positions.DELETE("/:id", server.deletePosition)
		positions.PUT("/:id", server.updatePosition)
	}

	teams := v1.Group("/teams")
	{
		teams.GET("/", server.getTeams)
		teams.GET("/:id", server.getTeamById)
		teams.POST("/", server.createTeam)
		teams.DELETE("/:id", server.deleteTeam)
		teams.PUT("/:id", server.updateTeam)
	}

	employees := v1.Group("/employees")
	{
		employees.GET("/", server.getEmployees)
		employees.GET("/:id", server.getEmployeeById)
		employees.POST("/")
		employees.DELETE("/:id", server.deleteEmployee)
		employees.PUT("/:id")
	}

	employees_project := v1.Group("/assignments")
	{
		// Assign an employee to a project
		employees_project.POST("/project/:id_project/employee/:id_employee", server.AddUserToProject)
		// Select all the employees in a project
		employees_project.GET("/project/:id/employees", server.getProjectEmployees)
		// Select all the projects assigned to a single employee
		employees_project.GET("/employee/:id/projects", server.getEmployeeProjects)
		// Delete an employee from a project
		employees_project.DELETE("/employee/:id_employee/project/:id_project", server.deleteEmployeeFromAProject)
	}

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
