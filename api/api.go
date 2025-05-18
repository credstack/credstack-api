package api

/*
API - Provides a core abstraction of the API itself. Primarily handles dependency injection and graceful termination
of the API
*/
type API struct {
}

/*
New - A constructor for the API
*/
func New() *API {
	return &API{}
}
