package main 
import (
	"github.com/urfave/cli/v2"
	"fmt"
	"log"
	"os"
)
func main() {
	app := &cli.App{ // initializes a new cli app using urfave/cli package.
		Name: "HealthChecker", //takes the name of the application.
		Usage: "A website health status checker", //description of what the app does.
		Flags: []cli.Flag{ // pass in flags you will use to call a function in the app.

			//the first flag is simply to call the domain of web server you trying to check.
			&cli.StringFlag{
			Name: "domain",
			Aliases: []string{"d"},
			Usage: "Domain name to check.",
			Required: true,
		},
		&cli.StringFlag{
			//same as the top but for the port number.
			Name: "port",
			Aliases: []string{"p"},
			Usage: "Port number to check.",
			Required: false,
		},
	},
	//this line represent actions to be taken when flags are being passed,
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == ""  {
				port = "80"
			}
			//call the check function that can be found in our check.go file.
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	//run the commands or cli arguements using the os package.
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}