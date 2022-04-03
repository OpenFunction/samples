package hook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/xanzy/go-gitlab"
)

func ServeWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received webhook: %s", r.URL.Path)

	event, err := parse(r)
	if err != nil {
		// return 500 error code
		log.Printf("Error parsing webhook: %s", err)
		http.Error(w, "Error parsing webhook", 500)
		return
	}

	if err := handle(event); err != nil {
		// return 500 error code
		log.Printf("Error handling webhook: %s", err)
		http.Error(w, "Error handling webhook", 500)
		return
	}

	// return 200 OK
	w.WriteHeader(http.StatusOK)
}

func handle(event interface{}) error {
	str, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("could not marshal json event for logging: %v", err)
	}

	// write the event for this example
	fmt.Println(string(str))

	return nil
}

// parse verifies and parses the events specified in the request and
// returns the parsed event or an error.
func parse(r *http.Request) (interface{}, error) {
	defer func() {
		if _, err := io.Copy(ioutil.Discard, r.Body); err != nil {
			log.Printf("Could discard request body: %v", err)
		}
		if err := r.Body.Close(); err != nil {
			log.Printf("Could not close request body: %v", err)
		}
	}()

	if r.Method != http.MethodPost {
		return nil, errors.New("invalid HTTP Method")
	}

	event := r.Header.Get("X-Gitlab-Event")
	if strings.TrimSpace(event) == "" {
		return nil, errors.New("missing X-Gitlab-Event Header")
	}

	eventType := gitlab.EventType(event)
	log.Printf("eventType: %s", eventType)

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, errors.New("error reading request body")
	}

	return gitlab.ParseWebhook(eventType, payload)
}
