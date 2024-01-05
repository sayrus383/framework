package models

type EndpointMapping struct {
	Id       int
	Source   string
	Target   string
	Disabled bool
	IsAuth   bool
}
