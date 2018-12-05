
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
}

func _pares_source() string {
    _, filename, line, _ := runtime.Caller(3)
    filename, _ = filepath.Rel(root, filename)
    return fmt.Sprintf("[%s:%d]", filename, line)
}

func _printf(format string, a ...interface{}) {
    logger.Printf("%s %s\n",_pares_source(), fmt.Sprintf(format, a...))
}

func _println(a ...interface{}) {
    tmp := []interface{}{_pares_source()}
    for _, val := range a {
        tmp = append(tmp, val)
    }

    logger.Println(tmp...)
}

func Debugf(format string, a ...interface{}) {
    if logLevel > 2 {
        _printf(format, a...)
    }
}

func Debug(a ...interface{}){
    if logLevel > 2 {
        _println(a...)
    }
}

func Infof(format string, a ...interface{}) {
    if logLevel > 1 {
        _printf(format, a...)
    }
}

func Info (a ...interface{}) {
    if logLevel > 1 {
        _println(a...)
    }
}


func Errorf (format string, a ...interface{}) {
    if logLevel > 0 {
        _printf(format, a...)
    }
}

func Error (a ...interface{}) {
    if logLevel > 0 {
        _println(a...)
    }
}


func Panic(format string, a ...interface{}) {
    _printf(format, a...)
    panic("!!")
}

func Logf(format string, a ...interface{}) {
    _printf(format, a...)
}

func Log(a ...interface{}) {
    _println(a...)
}

func Stack(format string, a ...interface{}){
    _printf(format, a...)
    buf := make([]byte, 8192)
    runtime.Stack(buf, false)
    _printf("!!!!!stack!!!!!: %s", buf)
}

func Recover() {
    if err := recover(); err != nil {
        Stack("goroutine failed:%v", err)
    }
}
