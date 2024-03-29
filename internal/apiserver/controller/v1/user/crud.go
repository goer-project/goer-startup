package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"goer-startup/internal/pkg/core"
	"goer-startup/internal/pkg/errno"
	"goer-startup/internal/pkg/log"
	v1 "goer-startup/pkg/api/goer/v1"
)

// List 返回用户列表，只有 root 用户才能获取用户列表.
//
// @Summary    List users
// @Security   Bearer
// @Tags       User
// @Accept     application/json
// @Produce    json
// @Param      request	 query	    v1.ListUserRequest	 true  "Param"
// @Success	   200		{object}	v1.ListUserResponse
// @Failure	   400		{object}	core.ErrResponse
// @Failure	   500		{object}	core.ErrResponse
// @Router    /v1/users [GET]
func (ctrl *UserController) List(c *gin.Context) {
	log.C(c).Infow("List user function called")

	var r v1.ListUserRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Users().List(c, r.Offset, r.Limit)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}

// Create 创建一个新的用户.
//
// @Summary    Create a user
// @Security   Bearer
// @Tags       User
// @Accept     application/json
// @Produce    json
// @Param      request	 body	    v1.CreateUserRequest	 true  "Param"
// @Success	   200		{object}	v1.GetUserResponse
// @Failure	   400		{object}	core.ErrResponse
// @Failure	   500		{object}	core.ErrResponse
// @Router    /v1/users [POST]
func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	// Validator
	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	// Create user
	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	// Create policy
	if _, err := ctrl.a.AddNamedPolicy("p", r.Username, "/v1/users/"+r.Username, defaultMethods); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}

// Get 获取一个用户的详细信息.
//
// @Summary    Get user info
// @Security   Bearer
// @Tags       User
// @Accept     application/json
// @Produce    json
// @Param      name	     path	    string          	 true  "Username"
// @Success	   200		{object}	v1.ListUserResponse
// @Failure	   400		{object}	core.ErrResponse
// @Failure	   500		{object}	core.ErrResponse
// @Router    /v1/users/{name} [GET]
func (ctrl *UserController) Get(c *gin.Context) {
	log.C(c).Infow("Get user function called")

	user, err := ctrl.b.Users().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}

// Update 更新用户信息.
//
// @Summary    Update user info
// @Security   Bearer
// @Tags       User
// @Accept     application/json
// @Produce    json
// @Param      name	     path	    string          	 true  "Username"
// @Param      request	 query	    v1.UpdateUserRequest	 true  "Param"
// @Success	   200		{object}	nil
// @Failure	   400		{object}	core.ErrResponse
// @Failure	   500		{object}	core.ErrResponse
// @Router    /v1/users/{name} [PUT]
func (ctrl *UserController) Update(c *gin.Context) {
	log.C(c).Infow("Update user function called")

	var r v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().Update(c, c.Param("name"), &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}

// Delete 删除一个用户.
//
// @Summary    Delete a user
// @Security   Bearer
// @Tags       User
// @Accept     application/json
// @Produce    json
// @Param      name	     path	    string          	 true  "Username"
// @Success	   200		{object}	nil
// @Failure	   400		{object}	core.ErrResponse
// @Failure	   500		{object}	core.ErrResponse
// @Router    /v1/users/{name} [DELETE]
func (ctrl *UserController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete user function called")

	username := c.Param("name")

	if err := ctrl.b.Users().Delete(c, username); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if _, err := ctrl.a.RemoveNamedPolicy("p", username, "", ""); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
