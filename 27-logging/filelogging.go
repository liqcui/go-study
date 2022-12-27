package main

type FileLogger struct {
	Level       Loglevel
	filePath    string
	fileName    string
	errFileName string
	maxFileSize int64
}

func NewFileLogger(levelStr,fp,fn string,maxSize int64) *FileLogger{
   logLevel,err:=parseLogLevel(levelStr)
   if err!=nil{
	  panic(err) 
   }

   return &FileLogger{
	   
   }

}