package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"govas/utils"
	"net/http"
)

func main() {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("Jindo_GoVASSSSSSSSSSSSSSSSSSS0912093809817238"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", utils.Login)
	r.GET("/logout", utils.Logout)
	private := r.Group("/")
	private.Use(AuthRequired())
	{
		//user
		private.GET("/users", utils.GetUsers)
		private.POST("/users", utils.CreateUser)
		private.GET("/users/:uuid", utils.GetUser)
		private.DELETE("/users", utils.DelUser)
		private.PUT("/users", utils.UpdateUser)

		//role
		private.GET("/roles", utils.GetRoles)
		private.GET("/roles/:uuid", utils.GetRole)

		//target
		private.GET("/targets/:trash/:tasks", utils.GetTargets)
		private.POST("/targets", utils.CreateTarget)
		private.GET("/target/:id", utils.GetTarget)
		private.DELETE("/targets/:ultimate/:id", utils.DeleteTarget)
		private.POST("/targets/:id/clone", utils.CloneTarget)
		private.PUT("/targets", utils.UpdateTarget)

		//task
		private.GET("/tasks", utils.GetTasks)
		private.GET("/tasks/:task_id/:action", utils.ActionTask)
		private.GET("/tasks/:task_id", utils.GetTask)
		private.DELETE("/tasks/:task_id/:ultimate", utils.DeleteTask)
		private.POST("/tasks", utils.CreateTask)
		private.PUT("/tasks", utils.UpdateTask)
		//report
		private.GET("/reports", utils.GetReports)
		private.GET("/reports/:report_id", utils.GetReport)
		private.DELETE("/reports/:report_id", utils.DeleteReport)
		//result
		private.GET("/results", utils.GetResults)
		private.GET("/results/:result_id", utils.GetResult)
		// schedule
		private.GET("/schedules", utils.GetSchedules)
		private.GET("/schedules/:schedule_id/clone", utils.CloneSchedule)
		private.GET("/schedules/:schedule_id", utils.GetSchedule)
		private.DELETE("/schedules/:schedule_id/:ultimate", utils.DeleteSchedule)
		private.DELETE("/schedules/:schedule_id", utils.DeleteSchedule)
		private.POST("/schedules", utils.CreateSchedule)
		private.PUT("/schedules", utils.UpdateSchedule)

	}
	_ = r.Run(":8080")
}
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		//session.Options(sessions.Options{MaxAge: 5})
		user := session.Get("username")
		//_ = session.Save()
		if user == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		} else {
			c.Next()
		}
	}
}
