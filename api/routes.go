package api

import (
	"gomboc/api/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secure      bool
	HandleFunc  http.HandlerFunc
	Description string
}

var routes = []Route{
	{
		Name:        "UserInfo",
		Method:      http.MethodGet,
		Pattern:     "/user/profile",
		Secure:      true,
		HandleFunc:  handlers.UserInfoHandler,
		Description: "Returns user information form id",
	},
	{
		Name:        "CreateUser",
		Method:      http.MethodPost,
		Pattern:     "/user",
		Secure:      true,
		HandleFunc:  handlers.CreateUserHandler,
		Description: "Create a user",
	},
	{
		Name:        "UpdateUser",
		Method:      http.MethodPatch,
		Pattern:     "/user",
		Secure:      true,
		HandleFunc:  handlers.UpdateUserHandler,
		Description: "Update a user",
	},
	{
		Name:        "DeleteUser",
		Method:      http.MethodDelete,
		Pattern:     "/user",
		Secure:      true,
		HandleFunc:  handlers.DeleteUserHandler,
		Description: "Delete user from database",
	},
	{
		Name:        "NodesInfo",
		Method:      http.MethodGet,
		Pattern:     "/nodes",
		Secure:      true,
		HandleFunc:  handlers.GetNodesHandler,
		Description: "Return a list of nodes with short information",
	},
	{
		Name:        "CreateNode",
		Method:      http.MethodPost,
		Pattern:     "/node",
		Secure:      true,
		HandleFunc:  handlers.CreateNodeHandler,
		Description: "Create a new node",
	},
	{
		Name:        "GrantNodeAccess",
		Method:      http.MethodPost,
		Pattern:     "/node/grant-access",
		Secure:      true,
		HandleFunc:  handlers.GrantNodeAccessHandler,
		Description: "Grant node permissions for a specific host",
	},
	{
		Name:        "RevokeNodeAccess",
		Method:      http.MethodPost,
		Pattern:     "/node/revoke-access",
		Secure:      true,
		HandleFunc:  handlers.RevokeNodeAccessHandler,
		Description: "Revoke node permissions for a specific host",
	},
	{
		Name:        "StartNodeSession",
		Method:      http.MethodPost,
		Pattern:     "/node/node-session",
		Secure:      false,
		HandleFunc:  handlers.StartNodeSessionHandler,
		Description: "Create a record with a node session",
	},
	{
		Name:        "RequestNodeRegistration",
		Method:      http.MethodPost,
		Pattern:     "/node/request-registration",
		Secure:      false,
		HandleFunc:  handlers.RequestNodeRegistrationHandler,
		Description: "Creates a device and returns host trusted info",
	},
}
