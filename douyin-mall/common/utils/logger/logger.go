// 日志记录
package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

var Log *zap.Logger

func Init() error {
	// 创建日志目录
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// 创建日志文件
	appLogPath := filepath.Join(logDir, "app.log")
	errorLogPath := filepath.Join(logDir, "error.log")

	// 确保日志文件存在
	if _, err := os.OpenFile(appLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); err != nil {
		return fmt.Errorf("failed to create app log file: %v", err)
	}
	if _, err := os.OpenFile(errorLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); err != nil {
		return fmt.Errorf("failed to create error log file: %v", err)
	}

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", appLogPath}
	config.ErrorOutputPaths = []string{"stderr", errorLogPath}

	var err error
	Log, err = config.Build()
	if err != nil {
		return fmt.Errorf("failed to build logger: %v", err)
	}

	Log.Info("Logger initialized successfully")
	return nil
}

// 使用示例：
// logger.Log.Info("payment processed",
//     zap.String("transaction_id", tx.ID),
//     zap.Float64("amount", amount))
