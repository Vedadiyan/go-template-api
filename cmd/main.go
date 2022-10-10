package main

import (
	autoconfix "booqall.com/libraries/autoconfix/pkg"
	runtmx "booqall.com/libraries/runtmx/pkg"
	core "booqall.com/libraries/webapix/pkg"
	"github.com/sirupsen/logrus"
)

func initializers() []autoconfix.Initializer {
	initializers := make([]autoconfix.Initializer, 0)

	// To Do

	return initializers
}

func router(baseRoute string) {
	route := core.NewRoute(baseRoute)
	_ = route
}

func init() {
	autoconfix.Bootstrap(initializers()...)
	router(core.ROOT)
}

func main() {
	engine := core.NewEngine()
	engine.DisableCORS()
	engine.Start("0.0.0.0:5000")
	logrus.Info("Server Started")
	runtmx.AwaitInterrupt(func() {
		logrus.Warn("Server is Shutting Down")

		logrus.Warn("Server Shut Down")
	})
}
