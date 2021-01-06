package httpcalls_test

import (
	"fmt"
	"log"
	"time"

	"github.com/marco-ostaska/httpcalls"
)

func ExampleAPIData_NewRequest() {
	httpcalls.Insecure = true // used for invalid certificates

	var facts []struct {
		Status struct {
			Verified  bool `json:"verified"`
			SentCount int  `json:"sentCount"`
		} `json:"status"`
		Type      string    `json:"type"`
		Deleted   bool      `json:"deleted"`
		ID        string    `json:"_id"`
		User      string    `json:"user"`
		Text      string    `json:"text"`
		V         int       `json:"__v"`
		Source    string    `json:"source"`
		UpdatedAt time.Time `json:"updatedAt"`
		CreatedAt time.Time `json:"createdAt"`
		Used      bool      `json:"used"`
	}

	var a httpcalls.APIData
	a.URL = "https://cat-fact.herokuapp.com"
	a.API = "/facts"
	if err := a.NewRequest(&facts); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(facts[0].Text)

	// Output:
	// Cats make about 100 different sounds. Dogs make only about 10.

}

func ExampleAPIData_QueryGraphQL() {

	var RickMorty struct {
		Data struct {
			EpisodesByIds []struct {
				Name       string `json:"name"`
				Characters []struct {
					Name string `json:"name"`
				} `json:"characters"`
			} `json:"episodesByIds"`
		} `json:"data"`
	}

	var a httpcalls.APIData
	a.URL = "https://rickandmortyapi.com/graphql/"
	query := `{
		episodesByIds(ids: [1, 2]) { 
			name characters {name}
		}
	}`

	if err := a.QueryGraphQL(query, &RickMorty); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(RickMorty.Data.EpisodesByIds[0].Characters[0].Name)

	// Output:
	// Rick Sanchez

}
