package types

type (
	User struct {
		Username string `json:"username"`
		AssignedRoles []string `json:"assignedRoles"`
		AuthorizedRoles []string `json:"authorizedRoles"`
	}

	Session struct {
		ID string `json:"session"`
		Username string `json:"username"`
		Roles []string `json:"roles"`
	}

	// @todo: need to list nested roles,
	// @todo: don't return users=null - return users: []?
	Role struct {
		Name string `json:"rolename"`
		Users []string `json:"users"`
		Permissions []string `json:"users"`
	}
)