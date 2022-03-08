# Functions framework demos

This directory holds the validation demos for the [functions-framework](https://github.com/OpenFunction/functions-framework).

Current supported list:

- [functions-framework-go](https://github.com/OpenFunction/functions-framework-go)

Depending on the type of runtime, the demos are divided into `Knative` and `Async`, please check them separately:

* Golang
  * [Knative runtime demos](golang/Knative)
  * [Async runtime demos](golang/Async)

## Plugin mechanism

> You can preferably refer to this [proposal](https://github.com/OpenFunction/OpenFunction/blob/main/docs/proposals/202112_functions_framework_refactoring.md) to learn more about how the function framework works.

Take the **functions-framework-go** as an example, we provide a default plugin called `plugin-example`, which is integrated inside the **functions-framework-go**.

There is also a plugin called `plugin-custom` which is a user-defined plugin.

When defining `FUNC_CONTEXT`, the user needs to configure the contents of `$.prePlugins` and `$.postPlugins` according to the execution order of the plugins.

In *demos*, we configure the order of plugins as follows.

```json
  "prePlugins": ["plugin-custom", "plugin-example"],
  "postPlugins": ["plugin-custom", "plugin-example"]
```

The processing in the plugin is as follows:

plugin-custom

```go
func (p *PluginCustom) ExecPreHook(ctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	p.stateC = 3
	p.StateC = 3
	return nil
}

func (p *PluginCustom) ExecPostHook(ctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	return nil
}
```

plugin-example

```go
func (p *PluginExample) ExecPreHook(ctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	r := preHookLogic(ctx.Ctx)
	p.stateA = 1
	p.stateB = r
	return nil
}

func (p *PluginExample) ExecPostHook(ctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	// Get data from another plugin via Plugin.Get()
	plgName := "plugin-custom"
	keyName := "StateC"
	plg, ok := plugins[plgName]
	if ok && plg != nil {
		v, exist := plg.Get(keyName)
		if exist {
			stateC := v.(int64)
			postHookLogic(p.stateA, stateC)
			return nil
		}
	}
	return fmt.Errorf("failed to get %s from plugin %s", keyName, plgName)
}

func preHookLogic(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	} else {
		return context.Background()
	}
}

func postHookLogic(numA int64, numB int64) int64 {
	sum := numA + numB
	klog.Infof("the sum is: %d", sum)
	return sum
}
```

As you can see, the two plugins, when combined, will print the following message:

```shell
the sum is: 4
```

