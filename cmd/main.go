package main

import (
	"github.com/alecthomas/kong"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/tanalam2411/cp-pitbox/cmd/local"
	"github.com/kyokomi/emoji/v2"
)


type cli struct {
	Local local.Cmd `cmd:"" help:"Local setup."`
}

func main() {
	logger := logging.NewNopLogger()
	pizzaMessage := emoji.Sprint(":flashlight:")
	ctx := kong.Parse(&cli{},
		kong.Name(pizzaMessage + "Crossplane PitBox "),
		kong.Description("Crossplane PitBox desc..."),
		kong.BindTo(logger, (*logging.Logger)(nil)),
		kong.ConfigureHelp(kong.HelpOptions{
			FlagsLast: true,
			Compact: true,
			WrapUpperBound: 80,
		}),
		kong.UsageOnError(),
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}