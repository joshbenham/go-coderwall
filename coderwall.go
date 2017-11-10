package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// CoderwallUser stores the user json
type CoderwallUser struct {
	ID       int
	Username string
}

// Coderwall stores the json response
type Coderwall struct {
	User []CoderwallUser
}

func main() {
	app := cli.NewApp()
	app.Name = "Coderwall"
	app.Usage = "Fetch Coderwall stats for a user"
	app.UsageText = "coderwall <username>"

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			return StyleError(app.UsageText)
		}

		username := c.Args().Get(0)
		url := "https://coderwall.com/" + username + ".json"

		StyleHeading("\n -> " + url + "\n")
		response, err := CallAPI(url)

		if err != nil {
			return StyleError(fmt.Sprintf("Could not Convert to Json: %s", err))
		}

		collection := response.(map[string]interface{})
		user := collection["user"].(map[string]interface{})

		fmt.Printf("ID: %v\n", user["id"])
		fmt.Printf("User: %v (%v)\n", user["name"], user["username"])
		fmt.Printf("Title: %v\n", user["title"])
		fmt.Printf("Endorsements: %v\n", user["endorsements"])

		badges := user["badges"].([]interface{})
		fmt.Printf("Badges: %v\n", len(badges))

		for _, element := range badges {
			badge := element.(map[string]interface{})
			fmt.Printf(" - %v (%v)\n", badge["name"], badge["description"])
		}

		return nil
	}

	app.Run(os.Args)
}
