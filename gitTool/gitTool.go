package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/levigross/grequests"
	"github.com/urfave/cli"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

// Struct for holding response of repo fetch API
type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullname"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

// Struct for modelling JSON body in create Gist
type File struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	File        map[string]File `json:"files"`
}

// Fetches the repo for the given github acct
func getStats(url string) *grequests.Response {
	res, err := grequests.Get(url, requestOptions)

	// You can modify the request  b passiing an optional RequestOption struct
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return res
}

// Read the files provided and creates Gist on github
func createGist(url string, args []string) *grequests.Response {
	// get first teo arguments
	description := args[0]

	// remaining arguments are file names with path
	var fileContents = make(map[string]File)
	for i := 0; i < len(args); i++ {
		dat, err := ioutil.ReadFile(args[i])
		if err != nil {
			log.Println("Please check the filenames. Absolute path (or)  same directory are alowed")
			return nil
		}
		var file File
		file.Content = string(dat)
		fileContents[args[i]] = file
	}

	var gist = Gist{Description: description, Public: true, File: fileContents}
	var postBody, _ = json.Marshal(gist)
	var requestOptions_copy = requestOptions

	// Add data to json field
	requestOptions_copy.JSON = string(postBody)

	// make a post request to github
	res, err := grequests.Post(url, requestOptions_copy)

	if err != nil {
		log.Println("Create request failed for Github Api")
	}
	return res
}

func main() {
	app := cli.NewApp()

	// define command for the client
	// define command for our client
	app.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Fetch the repo details with user. [Usage]: goTool fetch user",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					// Github Api logic
					var repos []Repo
					username := c.Args()[0]
					var repoUrl = fmt.Sprintf("https://api.github.com/users/%s/repos", username)
					resp := getStats(repoUrl)
					resp.JSON(&repos)
					log.Println(repos)
				} else {
					log.Println("Please give a username. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "Create",
			Aliases: []string{"c"},
			Usage:   "Create a giste from the given test. [Usage]: goTool name 'description' sample.txt",
			Action: func(c *cli.Context) error {
				if c.NArg() > 1 {
					// Github Api Logic
					args := c.Args()
					var postUrl = "https://api.github.com/gists"
					resp := createGist(postUrl, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
	}
	app.Version = "1.0"
	app.Run(os.Args)

}
