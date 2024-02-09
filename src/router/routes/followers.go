package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var FollowersRoutes = []Route{
	{
		URI:          "/followers/all/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.RetrieveAllFollowers,
		AuthRequired: true,
	},
	{
		URI:          "/followers/following/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.RetrieveUsersFollowing,
		AuthRequired: true,
	},
	{
		URI:          "/followers/follow/{userId}",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/followers/unfollow/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.UnfollowUser,
		AuthRequired: true,
	},
}
