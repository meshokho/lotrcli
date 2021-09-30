package cmd

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/spf13/cobra"
	"math/rand"
	"time"
)

// phraseCmd represents the phrase command
var phraseCmd = &cobra.Command{
	Use:   "phrase",
	Short: "Get a random phrase",
	Long: `This command fetches a random phrase from LOTR using The One API`,
	Run: func(cmd *cobra.Command, args []string) {
		getPhrase()
	},
}


func init() {
	rootCmd.AddCommand(phraseCmd)
}


type urlParams struct {
	method 		string
	id			int 			
	submethod 	string
}

type Token struct {
	Token string `xml:"token"`
}


// Prints random phrase from LOTR.
func getPhrase() {
	rand.Seed(time.Now().UnixNano())

	var urlP urlParams
	urlP.method = "quote"

	res := connectToAPI(urlP)
	fmt.Println(res["docs"].([]interface{})[rand.Intn(1000)].(map[string]interface{})["dialog"].(string))
}


// Connects to API, returns JSON decoded responce body.
func connectToAPI(urlP urlParams) map[string]interface{} {
/*	xmlFile, err := ioutil.ReadFile("cmd/settings.xml")
	if err != nil {
		log.Fatal(err)
	}
	token := &Token{}
	xml.Unmarshal([]byte(xmlFile), &token)*/

	token := "ESYNktU1WlxM9l_6lMJ5"

	baseURL := "https://the-one-api.dev/v2"
	url := baseURL + "/" + urlP.method

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer " + token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
