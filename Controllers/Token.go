package controllers

type TokenController interface {
	index(*gin.Context)
}
type tokenController struct {