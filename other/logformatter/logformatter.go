package logformatter

import (
	"github.com/gyozatech/noodlog"
)

// @return noodlog.Logger
func New() *noodlog.Logger {
	return noodlog.NewLogger().SetConfigs(noodlog.Configs{
		LogLevel:             noodlog.LevelTrace,
		JSONPrettyPrint:      noodlog.Enable,
		TraceCaller:          noodlog.Enable,
		Colors:               noodlog.Enable,
		CustomColors:         &noodlog.CustomColors{Trace: noodlog.Cyan, Error: noodlog.Red, Warn: noodlog.Yellow},
		ObscureSensitiveData: noodlog.Enable,
		SensitiveParams:      []string{""},
	})
}
