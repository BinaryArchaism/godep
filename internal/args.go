package internal

import (
	"fmt"
	"time"
)

func ParseArgs(args []string) {
	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		switch arg {
		case "version":
			VersionCmd()
		case "help":
			helpCmd()
		default:
			unknownCmd(arg)
		}
	}
}

func unknownCmd(cmd string) {
	fmt.Printf("godep: unknown command: %s\n", cmd)
	fmt.Println("Run 'godep help' to see usage")
}

func helpCmd() {
	fmt.Println("GoDep help")
	fmt.Println("\tversion - prints version of GoDep tool")
	fmt.Println("\thelp - prints this sheet")
}

func VersionCmd() {
	fmt.Println(logoMsg)
	fmt.Println(time.Now().Format(time.RFC1123))
	fmt.Printf("Version: %s\n", CurrentVesion)
}

var logoMsg = `
 _____      ______           
|  __ \     |  _  \          
| |  \/ ___ | | | |___ _ __  
| | __ / _ \| | | / _ \ '_ \ 
| |_\ \ (_) | |/ /  __/ |_) |
 \____/\___/|___/ \___| .__/ 
                      | |    
                      |_|    
`
