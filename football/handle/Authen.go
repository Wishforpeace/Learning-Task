package handle

import (
	"Mygo/football/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func authenticationsHandler(w http.ResponseWriter, req *http.Request) {

	username, password, ok := req.BasicAuth()

	if ok {

		tokenDetails, err := model.generateToken(username, password)

		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {

			enc := json.NewEncoder(w)
			enc.SetIndent("", "  ")
			enc.Encode(tokenDetails)
		}
	} else {

		fmt.Fprintf(w, "You require a username/password to get a token.\r\n")
	}

}
