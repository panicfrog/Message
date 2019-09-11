package api

type ApiStatus int

const (
	ApiStatusSuccess            ApiStatus = 0
	ApiStatusFailed             ApiStatus = 1
	ApiStatusParamsError        ApiStatus = 2
	ApiStatusInternelError      ApiStatus = 3
	ApiStatusUnauthUnauthorized ApiStatus = 4
	ApiStatusTokenUnknowErr     ApiStatus = 5
)