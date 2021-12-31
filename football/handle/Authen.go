package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/model"
)

func AuthenticationsHandler(w http.ResponseWriter, req *http.Request) {

	username, password, ok := req.BasicAuth()

	if ok {

		tokenDetails, err := model.GenerateToken(username, password)

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
