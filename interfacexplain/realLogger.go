package interfacexplain

import "fmt"

type Logger interface{
	PrintMessage(message string) 
}

type FileLogging struct{}

func (f FileLogging) PrintMessage(message string) {
	fmt.Println("Logging for file", message)
}

type CloudLogging struct{}

func (c CloudLogging) PrintMessage(message string){
	fmt.Println("cloud logger", message)
}

func callLogging(logger Logger,message string){
   logger.PrintMessage(message)
}

func OwnLoggerForDiffrentType(){
	f := FileLogging{}
	callLogging(f,"savan")
	c := CloudLogging{}
	callLogging(c,"cloud")
}
