package quotes_controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quotes/src/quotes_db"
)

func NextQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Trying to fetch next quote")
	quotesList := quotes_db.GetQuotes("quotes")

	key := quotes_db.QuotesMapRandomKeyGet(quotesList).(string)
	quote := `{"quote":"` + quotesList[key] + `"}`

	jData, err := json.Marshal(quote)
	if err != nil {
		// handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jData))
}
