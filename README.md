# dynamic-level-handler

Golang wrapper for slog.Handler that supports level override for concrete logger. This functionality was originally implemented in [Neptunus](https://github.com/gekatateam/neptunus) to support different logging levels of pipelines and plugins.

## But... Why?
Have you ever wanted to enable only warnings or errors for specific parts of an application? And then enable debug for those concrete parts? I have. This package was created with that idea in mind.

`DynamicHandlerWrapper` wraps the passed `slog.Handler` and allows you to set a personal logging level via `slog.Leveler`. If level has not been overridden, all calls are passed to the basic handler.

## Examples
```go
import (
	"log/slog"

	dynamic "github.com/gekatateam/neptunus/dynamic-level-handler"
)

// create basic handler
h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo
})

// wrap basic handler
dynamicHandler := dynamic.New(h)

// create common logger for all app
logger := slog.New(dynamicHandler)

......

// create child logger with some attributes
// then override it's level
pipelineLogger := logger.With(slog.Group("pipeline",
    "id", "cronjobs",
))
// you don't need to think about what specific handler type is passed here
// this function only does something if handler type is `DynamicLevelHandler`
dynamic.OverrideLevel(pipelineLogger.Handler(), slog.LevelWarn)

......

// create another child logger, set it's personal level
pluginLogger := pipelineLogger.log.With(slog.Group("input",
    "plugin", "cronjob",
    "name", "first cron",
))
dynamic.OverrideLevel(pluginLogger.Handler(), slog.LevelDebug)
```
