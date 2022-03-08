package main

import (
	"context"

	"github.com/OpenFunction/functions-framework-go/framework"
	"github.com/OpenFunction/functions-framework-go/plugin"
	"k8s.io/klog/v2"
	"main.go/userfunction"
)

func main() {
	ctx := context.Background()
	fwk, err := framework.NewFramework()
	if err != nil {
		klog.Exit(err)
	}
	fwk.RegisterPlugins(getLocalPlugins())
	if err := fwk.Register(ctx, userfunction.Subscriber); err != nil {
		klog.Exit(err)
	}
	if err := fwk.Start(ctx); err != nil {
		klog.Exit(err)
	}
}

func getLocalPlugins() map[string]plugin.Plugin {
	localPlugins := map[string]plugin.Plugin{}

	if len(localPlugins) == 0 {
		return nil
	} else {
		return localPlugins
	}
}
