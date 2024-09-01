package contracts

import (
	"github.com/LuanTruongPTIT/go-food/internal/pkg/config/environment"
	"github.com/LuanTruongPTIT/go-food/internal/pkg/logger"
	"go.uber.org/fx"
)

type ApplicationBuilder interface {
	ProvideModule(module fx.Option)
	Provide(constructors ...interface{})
	Decorate(constructors ...interface{})
	Build() Application
	GetProvides() []interface{}
	GetDecorates() []interface{}
	Options() []fx.Option
	Logger() logger.Logger
	Environment() environment.Environment
}
