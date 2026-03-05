package main

import (
	"log/slog"
	"time"

	"github.com/go-drift/drift/pkg/core"
	"github.com/go-drift/drift/pkg/drift"
	"github.com/go-drift/drift/pkg/graphics"
	"github.com/go-drift/drift/pkg/layout"
	"github.com/go-drift/drift/pkg/overlay"
	"github.com/go-drift/drift/pkg/theme"
	"github.com/go-drift/drift/pkg/widgets"
)

func main() {
	drift.NewApp(App()).Run()
}

func App() core.Widget {
	// https://driftframework.dev/docs/guides/theming#providing-theme
	return theme.Theme{
		Data: theme.DefaultDarkTheme(), // or DefaultLightTheme()
		// We add an overlay to use later with Toast
		Child: overlay.Overlay{
			Child: app{},
		},
	}
}

type app struct {
	core.StatefulBase
}

func (app) CreateState() core.State {
	return &appState{}
}

// https://driftframework.dev/docs/guides/state-management#lifecycle-order
type appState struct {
	core.StateBase
}

func (s *appState) InitState() {
}

func (s *appState) Build(ctx core.BuildContext) core.Widget {
	_, colors, _ := theme.UseTheme(ctx)

	return widgets.Container{
		Color: colors.Background,
		Child: widgets.Centered(
			theme.ButtonOf(ctx, "Click me", func() {
				Toast(ctx, "Thank you for clicking")
				// tts := tooltipState{}
				// tts.showTooltip(ctx, "Thank you :)")
			}),
		),
	}
}

func Toast(ctx core.BuildContext, message string) {
	overlayState := overlay.OverlayOf(ctx)
	if overlayState == nil {
		slog.Error("Toast: Unable to get overlay to display Toast")
		return
	}

	colors := theme.ColorsOf(ctx)

	var entry *overlay.OverlayEntry

	entry = overlay.NewOverlayEntry(func(ctx core.BuildContext) core.Widget {
		// Use a Stack that fills the overlay to properly position the toast
		return widgets.Stack{
			Fit: widgets.StackFitExpand,
			Children: []core.Widget{
				widgets.Positioned(widgets.Container{
					Color:        colors.InverseSurface,
					BorderRadius: 8,
					Padding:      layout.EdgeInsetsSymmetric(20, 12),
					Child: widgets.Text{
						Content: message,
						Style: graphics.TextStyle{
							Color:      colors.OnInverseSurface,
							FontSize:   14,
							FontWeight: graphics.FontWeightBold,
						},
					},
				}).Align(graphics.AlignBottomCenter).Bottom(100),
			},
		}
	})

	overlayState.Insert(entry, nil, nil)

	go func() {
		time.Sleep(3 * time.Second)
		drift.Dispatch(entry.Remove)
	}()
}

type tooltipState struct {
	core.StateBase
	entry *overlay.OverlayEntry
}

func (s *tooltipState) showTooltip(ctx core.BuildContext, message string) {
	overlayState := overlay.OverlayOf(ctx)
	if overlayState == nil {
		slog.Error("Unable to get OverlayOf()")
		return
	}

	s.entry = overlay.NewOverlayEntry(func(ctx core.BuildContext) core.Widget {
		return widgets.Positioned(widgets.Container{
			Padding: layout.EdgeInsetsAll(8),
			Color:   graphics.RGBA(50, 50, 50, 0.9),
			Child:   widgets.Text{Content: message},
		}).At(100, 200)
	})
	overlayState.Insert(s.entry, nil, nil)
}

func (s *tooltipState) hideTooltip() {
	if s.entry != nil {
		s.entry.Remove()
		s.entry = nil
	}
}

func (s *tooltipState) Dispose() {
	s.hideTooltip()
	s.StateBase.Dispose()
}
