package logo

import (
	"fmt"
	"runtime"
	"time"

	"github.com/twistingmercury/go-figure"
)

func PrintConsoleLogo() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Cyan = "\033[36m"

	myFigure := figure.NewColorFigure("Gaggle", "larry3d", "red", true)
	myFigure.Print()
	fmt.Printf("\n%s                        %s\n", Red+">Authors:"+Reset, "Nora Cam & Nicholas Sutton")
	fmt.Printf("%s                      MIT License\n", Magenta+">Copyright:"+Reset)
	fmt.Printf("%s                           %s\n", Blue+">Time:"+Reset, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%s                             %s\n", Cyan+">OS:"+Reset, runtime.GOOS)
	fmt.Printf("%s                      Go & JavaScript\n", Green+">Languages:"+Reset)
}
