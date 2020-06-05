package main

import (
	"fmt"

	"github.com/orderforme/user/config"
	"github.com/orderforme/user/followers/follow"
	"github.com/orderforme/user/middelware"
	"github.com/orderforme/user/public/public"
	"github.com/orderforme/user/router"
	"github.com/orderforme/user/user/user"
)

func main() {
	// create a new echo instance
	r := router.New()

	route := r.Group("/api")

	middelware.Middelware(route)

	handlerUser, err := user.NewHandler()
	if err != nil {
		panic(err)
	}
	defer handlerUser.Done()

	handlerPublic, err := public.NewHandler()
	if err != nil {
		panic(err)
	}
	defer handlerPublic.Done()

	handlerFollow, err := follow.NewHandler()
	if err != nil {
		panic(err)
	}
	defer handlerPublic.Done()

	handlerUser.Register(route)

	middelware.MiddelwareJWT(route)

	handlerUser.RegisterJWT(route)
	handlerFollow.RegisterJWT(route)
	handlerPublic.RegisterJWT(route)

	r.Logger.Fatal(r.Start(fmt.Sprint(":", config.AppConfig.ServerPort)))
}
