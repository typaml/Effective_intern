package controllers

import (
	"Effective_intern/models"
	"Effective_intern/services"
	"Effective_intern/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: *services.NewUserService(),
	}
}

// @Summary Get users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	utils.Info("GetUsers: called")
	users, err := uc.userService.GetUsers()
	if err != nil {
		utils.Error("GetUsers: error getting users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("GetUsers: successfully retrieved users")
	c.JSON(http.StatusOK, users)
}

// @Summary Get user
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {array} models.User
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error("GetUserByID: invalid user ID:", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	utils.Info("GetUserByID: called with ID:", id)
	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		utils.Error("GetUserByID: error getting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("GetUserByID: successfully retrieved user with ID:", id)
	c.JSON(http.StatusOK, user)
}

// @Summary Get tasks
// @Description Get user tasks
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {array} models.User
// @Router /users/{id}/tasks [get]
func (uc *UserController) GetUserTasks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error("GetUserTasks: invalid user ID:", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	utils.Info("GetUserTasks: called with ID:", id)
	tasks, err := uc.userService.GetUserTasks(uint(id))
	if err != nil {
		utils.Error("GetUserTasks: error getting tasks:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("GetUserTasks: successfully retrieved tasks for user ID:", id)
	c.JSON(http.StatusOK, tasks)
}

// @Summary Create user
// @Description Create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param passportNumber body models.UnfetchingData true "passportNumber"
// @Success 200 {array} models.User
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var request models.UnfetchingData
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.Error("CreateUser: error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passportParts := strings.Split(request.PassportNumber, " ")
	if len(passportParts) != 2 {
		utils.Error("CreateUser: invalid passport format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid passport format"})
		return
	}

	passportSerie := passportParts[0]
	passportNumber := passportParts[1]
	user, err := utils.FetchPeopleInfo(passportSerie, passportNumber)
	if err != nil {
		utils.Error("CreateUser: error fetching people info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching people info"})
		return
	}
	if err := uc.userService.CreateUser(user); err != nil {
		utils.Error("CreateUser: error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("CreateUser: successfully created user")
	c.JSON(http.StatusCreated, *user)
}

// @Summary Update user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error("UpdateUser: invalid user ID:", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error("UpdateUser: error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.Info("UpdateUser: called with ID:", id)
	if err := uc.userService.UpdateUser(uint(id), &user); err != nil {
		utils.Error("UpdateUser: error updating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("UpdateUser: successfully updated user with ID:", id)
	c.JSON(http.StatusOK, user)
}

// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Error("DeleteUser: invalid user ID:", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	utils.Info("DeleteUser: called with ID:", id)
	if err := uc.userService.DeleteUser(uint(id)); err != nil {
		utils.Error("DeleteUser: error deleting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("DeleteUser: successfully deleted user with ID:", id)
	c.Status(http.StatusNoContent)
}

// @Summary Start tasks
// @Description Start user tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param taskId path int true "Task ID"
// @Success 200 {array} models.Task
// @Router /users/tasks/{id}/start [post]
func (uc *UserController) StartTask(c *gin.Context) {
	utils.Info("StartTask: called")
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		utils.Error("StartTask: invalid user ID:", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	utils.Info("StartTask: called with ID:", taskId)
	if err := uc.userService.StartTask(uint(taskId)); err != nil {
		utils.Error("DeleteUser: error deleting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("Start task: successfully start task user with ID:", taskId)
	c.Status(http.StatusOK)
}

// @Summary Stop tasks
// @Description Stop user tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param taskId path int true "Task ID"
// @Success 200 {array} models.Task
// @Router /users/tasks/{id}/stop [post]
func (uc *UserController) StopTask(c *gin.Context) {
	utils.Info("StopTask: called")
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		utils.Error("StartTask: invalid user ID:", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	utils.Info("StopTask: called with ID:", taskId)
	if err := uc.userService.StopTask(uint(taskId)); err != nil {
		utils.Error("DeleteUser: error deleting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.Info("Stop task: successfully stop task user with ID:", taskId)
	c.Status(http.StatusNoContent)
}
