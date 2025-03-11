package routes

import (
	"totesbackend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterItemTypeRoutes(router *gin.Engine, controller *controllers.ItemTypeController) {
	router.GET("/item-type", controller.GetItemTypes)
	router.GET("/item-type/:id", controller.GetItemTypeByID)
}

func RegisterItemRoutes(router *gin.Engine, controller *controllers.ItemController) {
	router.GET("/item/:id", controller.GetItemByID)
	router.GET("/item", controller.GetAllItems)
	router.GET("/item/searchById", controller.SearchItemsByID)
	router.GET("/item/searchByName", controller.SearchItemsByName)
	router.PATCH("/item/:id/state", controller.UpdateItemState)
	router.PUT("/item/:id", controller.UpdateItem)
}

func RegisterPermissionRoutes(router *gin.Engine, controller *controllers.PermissionController) {
	router.GET("/permission", controller.GetAllPermissions)
	router.GET("/permission/:id", controller.GetPermissionByID)
}

// //
func RegisterRoleRoutes(router *gin.Engine, controller *controllers.RoleController) {
	router.GET("/roles/:id", controller.GetRoleByID)
	router.GET("/roles/:id/permissions", controller.GetAllPermissionsOfRole)
	router.GET("/roles/:id/exists", controller.ExistRole)
	router.GET("/roles/", controller.GetAllRoles)
}

func RegisterUserTypeRoutes(router *gin.Engine, controller *controllers.UserTypeController) {
	router.GET("/user-types", controller.ObtainAllUserTypes)
	router.GET("/user-types/:id", controller.ObtainUserTypeByID)
	router.GET("/user-types/:id/exists", controller.Exists)
}

func RegisterUserStateTypeRoutes(router *gin.Engine, controller *controllers.UserStateTypeController) {
	router.GET("/user-state-type", controller.GetUserStateTypes)
	router.GET("/user-state-type/:id", controller.GetUserStateTypeByID)
}

func RegisterIdentifierTypeRoutes(router *gin.Engine, controller *controllers.IdentifierTypeController) {
	router.GET("/identifier-type", controller.GetIdentifierTypes)
	router.GET("/identifier-type/:id", controller.GetIdentifierTypeByID)
}

func RegisterUserRoutes(router *gin.Engine, controller *controllers.UserController) {
	router.GET("/user", controller.GetAllUsers)
	router.GET("/user/:id", controller.GetUserByID)
	router.GET("/user/searchByID", controller.SearchUsersByID)
	router.GET("/user/searchByEmail", controller.SearchUsersByEmail)
	router.PATCH("/user/:id/state", controller.UpdateUserState)
	router.PUT("/user/:id", controller.UpdateUser)
	router.POST("/user", controller.CreateUser)
}

func RegisterEmployeeRoutes(router *gin.Engine, controller *controllers.EmployeeController) {
	router.GET("/employee/", controller.GetAllEmployees)
	router.GET("/employee/:id", controller.GetEmployeeByID)
	router.GET("/employee/searchEmployeesByName", controller.SearchEmployeesByName)
	router.POST("/employee/", controller.CreateEmployee)
	router.PUT("/employee/:id", controller.UpdateEmployee)
	router.DELETE("/employee/:id", controller.DeleteEmployee)
}
