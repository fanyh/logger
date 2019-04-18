
package logger

import (
    "io"
    "log"
    "os"
    "flag"
    "runtime"
    "fmt"
    "path/filepath"
)

var logger *log.Logger
var logLevel int
var root string

func init() {
    logger = log.New(io.Writer(os.Stderr), "", log.Ldate|log.Lmicroseconds)
    flag.IntVar(&logLevel, "log", 2, "larger value for detail log")

    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        logger.Println("not read root dir")
        return
    }

    root = dir
    if filepath.Base(dir) == "bin" {
        root += "/.."
    }

    logger.Println("root patch", root)
}

func _pares_source(log_level string) string {
    _, filename, line, _ := runtime.Caller(3)
    filename, _ = filepath.Rel(root, filename)
    return fmt.Sprintf("[%s @./%s:%d]", log_level, filename, line)
}

func _printf(log_level string, format string, a ...interface{}) {
    logger.Printf("%s %s\n",_pares_source(log_level), fmt.Sprintf(format, a...))
}

func _println(log string, a ...interface{}) {
    tmp := []interface{}{_pares_source(log)}
    for _, val := range a {
        tmp = append(tmp, val)
    }

    logger.Println(tmp...)
}

func Debugf(format string, a ...interface{}) {
    if logLevel > 2 {
        _printf("DBG", format, a...)
    }
}

func Debug(a ...interface{}){
    if logLevel > 2 {
        _println("DBG", a...)
    }
}

func Infof(format string, a ...interface{}) {
    if logLevel > 1 {
        _printf("INF", format, a...)
    }
}

func Info (a ...interface{}) {
    if logLevel > 1 {
        _println("INF", a...)
    }
}


func Errorf (format string, a ...interface{}) {
    if logLevel > 0 {
        _printf("ERR",format, a...)
    }
}

func Error (a ...interface{}) {
    if logLevel > 0 {
        _println("ERR",a...)
    }
}


func Panic(format string, a ...interface{}) {
    _printf("DBG",format, a...)
    panic("!!")
}

func Logf(format string, a ...interface{}) {
    _printf("INF",format, a...)
}

func Log(a ...interface{}) {
    _println("INF",a...)
}

func Stack(format string, a ...interface{}){
    _printf("DBG",format, a...)
    buf := make([]byte, 8192)
    runtime.Stack(buf, false)
    _printf("DBG","!!!!!stack!!!!!: %s", buf)
}

func Recover() {
    if err := recover(); err != nil {
        Stack("goroutine failed:%v", err)
    }
}
