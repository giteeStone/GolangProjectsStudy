/*日志库
需求分析：
1.支持往不同的地方输出日志
2.日志分级别
  1.debug
  2.Trace
  3.info
  4.warning
  5.Error
  6.Fatal
3.日志要支持开关控制,比如开发的时候什么级别都能输出，但是上线之后只有INFO级别往下的才能输出
4.完整的日志记录要有时间、行号、文件名、日志级别、日志信息
5.日志文件要切割
  1.按文件大小切割
     1.每次记录日志之前都判断一下当前写的这个文件的文件大小
  2.按文件日期切割**/

package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

//往终端写日志相关内容

// Logger日志结构体
type LogLevel uint16

const (
	//定义日志级别
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// 根据string得到日志级别uint16
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}

// 由uint16日志级别，记录字符串string
func getlogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// 文件名
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

type Logger struct {
	Level LogLevel
}

// NewLog构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

// 要记录的日志file是否超过最大容量f.maxFileSize
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	//需要切割日志文件
	if f.checkSize(f.fileObj) {
		newFile, err := f.SplitFile(f.fileObj)
		if err != nil {
			fmt.Printf("split file error%v\n", err)
			return
		}
		f.fileObj = newFile
	}
	fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)
	if lv >= ERROR {
		if f.checkSize(f.errFileObj) {
			newFile, err := f.SplitFile(f.errFileObj)
			if err != nil {
				fmt.Printf("split err file error%v\n", err)
				return
			}
			f.errFileObj = newFile
			//如果要记录的日志大于等于ERROR级别，我还要在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

// 太大的文件切割保存日志
func (f *FileLogger) SplitFile(file *os.File) (*os.File, error) {
	//1.备份一下rename  xx.log->xx.log.bak20220406
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      //拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //拼接一个日志文件备份的名字
	//2.关闭当前的日志文件
	file.Close()

	os.Rename(logName, newLogName) //将原来的文件改名
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err%v\n", err)
		return nil, err
	}
	//3.将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, err
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		f.log(DEBUG, format, a...)
	}
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		f.log(INFO, format, a...)
	}
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.enable(WARNING) {
		f.log(WARNING, format, a...)
	}
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		f.log(ERROR, format, a...)
	}
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAL) {
		f.log(FATAL, format, a...)
	}
}

//往文件里写日志相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string //日志文件保存的路径
	fileName    string //日志文件保存的文件名
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

// 生产日志文件，
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	loglevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       loglevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile() //按照文件名和文件路径将文件打开
	if err != nil {
		panic(err)
	}
	return f1
}

// 打开日志文件备用
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}
	//日志文件都打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// 关闭日志文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func main() {
	log := NewFileLogger("ERROR", "./", "my.log", 10*100)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条Warning日志")
		// log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志,id:%d,name:%s", id, name)
		time.Sleep(2 * time.Second)
	}
}
