package main

func main() {

}

//
//import (
//	"./fetch"
//	"./logging"
//	"./slice"
//	"encoding/json"
//	"fmt"
//	"github.com/sirupsen/logrus"
//	"os"
//)
//
//func main() {
//	if hasHelpArg() {
//		printHelp()
//	}
//	logger := logging.GetInstance()
//	logger.SetLevel(logrus.DebugLevel)
//
//	requestParams := fetch.FetchRequestParams{Url: "http://www.google.com.br", Count: 2}
//	result, err := fetch.FetchRequest(requestParams)
//	if err != nil {
//		logger.Panic(err)
//	}
//
//	data, _ := json.Marshal(result)
//	logger.Printf(string(data))
//}
//
//func hasHelpArg() bool {
//	return utils.AnyInSlice(os.Args, func(s string) bool { return s == "help" })
//}
//
//func printHelp() {
//	fmt.Println("Usage: gourl <url> <fetchCount> [options]\n")
//	fmt.Println("Arguments:")
//	fmt.Println("<url>        The URL that you want to fetch")
//	fmt.Println("<fetchCount> Number of times that you want to fetch the URL (Default value is '1')\n")
//	fmt.Println("Options:")
//	fmt.Println("-v, --verbose  Enables verbose logging")
//	os.Exit(0)
//}
