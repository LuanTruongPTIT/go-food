package fxapp

import (
	"github.com/LuanTruongPTIT/go-food/internal/pkg/config/environment"
	"github.com/LuanTruongPTIT/go-food/internal/pkg/logger"
	"go.uber.org/fx"
)

type applicationBuilder struct {
	provides    []interface{}
	decorates   []interface{}
	options     []fx.Option
	logger      logger.Logger
	environment environment.Environment
}

// func NewApplicationBuilder(environments ...environment.Environment) contracts.ApplicationBuilder {
// 	env := environment.ConfigAppEnv(environments...)

// 	var logger logger.Logger
// 	lo
// }
