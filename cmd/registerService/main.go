package main

import (
	"context"
	"fmt"
	"github.com/cxb116/ADX_ENGINE/registerServer/registration"
	"github.com/cxb116/ADX_ENGINE/sspEngine/sspService"
)

func main() {

	ctx, err := registration.ServerEngineStart(context.Background(),
		"SspEngine Service",
		"127.0.0.1",
		"8080",
		sspService.RegisterHandler,
	)
	if err != nil {
		fmt.Println("SspEngine start error:", err)
	}
	fmt.Println("SspEngine start success")
	<-ctx.Done()

	fmt.Println("Shutting down SspEngine service .")
}
