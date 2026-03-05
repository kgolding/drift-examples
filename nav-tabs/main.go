package main

import (
	"log/slog"

	"github.com/go-drift/drift/pkg/core"
	"github.com/go-drift/drift/pkg/drift"
	"github.com/go-drift/drift/pkg/navigation"
	"github.com/go-drift/drift/pkg/theme"
	"github.com/go-drift/drift/pkg/widgets"
)

func main() {
	drift.NewApp(App()).Run()
}

func App() core.Widget {
	slog.Info("App()")

	tabController := navigation.NewTabController(0)

	// HACK to get it partly working!
	home := HomePage{}
	home.InitState()

	// https://driftframework.dev/docs/guides/theming#providing-theme
	return theme.Theme{
		Data: theme.DefaultDarkTheme(),
		Child: navigation.TabNavigator{
			Controller: tabController,
			Tabs: []navigation.Tab{
				navigation.NewTab(
					widgets.TabItem{
						Label: "Home",
						Icon:  HomeIcon(),
					},
					func(ctx core.BuildContext) core.Widget {
						return home.Build(ctx)
					},
				),
				navigation.NewTab(
					widgets.TabItem{
						Label: "Settings",
						Icon:  GearIcon(),
					},
					func(ctx core.BuildContext) core.Widget {
						page := SettingsPage{}
						return page.Build(ctx)
					},
				),
			},
		},
	}
}
