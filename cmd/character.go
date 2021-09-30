package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"time"
	"strings"
)

// characterCmd represents the character command
var characterCmd = &cobra.Command{
	Use:   "character",
	Short: "Get random character info",
	Long: `This command fetches a random phrase from LOTR using The One API`,
	Run: func(cmd *cobra.Command, args []string) {
		charName, _ := cmd.Flags().GetString("name")
		if charName == "" {
			charName = "Gandalf"
		}
		getCharacter(charName)
	},
}

func init() {
	rootCmd.AddCommand(characterCmd)

	characterCmd.PersistentFlags().String("name", "", "Character's name")
}

type Character struct {
	Id 		string
	Name 	string
	Birth 	string
	Death 	string
	Gender 	string
	Hair 	string
	Height 	string
	Race 	string
	Realm 	string
	Spouse 	string
	Wiki	string
}

func getCharacter(name string) {
	rand.Seed(time.Now().UnixNano())

	name = strings.ToLower(name)

	var urlP urlParams
	urlP.method = "character"

	res := connectToAPI(urlP)
	characters := res["docs"].([]interface{})

	var charDatas []Character

	for _, character := range characters {
		char := character.(map[string]interface{})
		if strings.Contains(strings.ToLower(char["name"].(string)), name) {
			charData := getCharData(char)
			charDatas = append(charDatas, charData)
		}
	}

	if len(charDatas) == 1 {
		fmt.Printf("Found 1 character:\n\n")
	} else if len(charDatas) > 1 {
		fmt.Printf("Found %d characters:\n\n", len(charDatas))
	}
	if len(charDatas) > 0 {
		for i, charData := range charDatas {
			printCharData(charData)
			if i != len(charDatas) - 1 {
				fmt.Println()
			}
		}		
	} else {
		fmt.Println("No character with this name found")
	}
}

func getCharData(char map[string]interface{}) Character {
	var newChar Character

	newChar.Id 		= getField("_id", char)
	newChar.Name 	= getField("name", char)
	newChar.Birth 	= getField("birth", char)
	newChar.Death 	= getField("death", char)
	newChar.Gender 	= getField("gender", char)
	newChar.Hair 	= getField("hair", char)
	newChar.Height 	= getField("height", char)
	newChar.Race 	= getField("race", char)
	newChar.Realm 	= getField("realm", char)
	newChar.Spouse 	= getField("spouse", char)
	newChar.Wiki	= getField("wikiUrl", char)

	return newChar
}

func printCharData(charData Character) {
	// fmt.Println("Id:	" 		+ charData.Id)
	fmt.Println("Name:	" 	+ charData.Name)
	fmt.Println("Birth:	" 	+ charData.Birth)
	fmt.Println("Death:	" 	+ charData.Death)
	fmt.Println("Race:	" 	+ charData.Race)
	fmt.Println("Gender:	" 	+ charData.Gender)
	fmt.Println("Hair:	" 	+ charData.Hair)
	fmt.Println("Height:	" 	+ charData.Height)
	fmt.Println("Race:	" 	+ charData.Race)
	fmt.Println("Realm:	" 	+ charData.Realm)
	fmt.Println("Spouse:	" 	+ charData.Spouse)
	fmt.Println("Wiki:	" 	+ charData.Wiki)	
}

func getField(field string, char map[string]interface{}) string {
	if _, ok := char[field].(string); !ok {
		return "No Data"
	} else {
		if char[field].(string) == "" {
			return "No Data"
		}
		return char[field].(string)
	}
}
