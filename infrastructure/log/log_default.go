package log

import (
  "context"
  "fmt"
)

// logPrinterDefault is default implementation of LogPrinter
type logPrinterDefault struct {
}

// WriteContext passing data to
func (r *logPrinterDefault) WriteContext(ctx context.Context, data ...interface{}) context.Context {
  return ctx
}

// LogPrint simply print the message to console
func (r *logPrinterDefault) LogPrint(ctx context.Context, flag string, data interface{}) {
  fmt.Printf(">>> %s\n", data)
}
