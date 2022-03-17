package logshandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	alert "github.com/prometheus/alertmanager/template"
)

const (
	HTTPCodeNotFound           = "404"
	Namespace                  = "demo-project"
	PodName                    = "wordpress-v1-[A-Za-z0-9]{5,15}-[A-Za-z0-9]{3,10}"
	AlertName                  = "404 Request"
	Severity                   = "warning"
	NotificationManagerAddress = "http://notification-manager-svc.kubesphere-monitoring-system.svc.cluster.local:19093/api/v2/alerts"
)

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Get request body failed")
		return
	}
	content := string(reqBody)
	matchHTTPCode, _ := regexp.MatchString(fmt.Sprintf(" %s ", HTTPCodeNotFound), content)
	matchNamespace, _ := regexp.MatchString(fmt.Sprintf("namespace_name\":\"%s", Namespace), content)
	matchPodName := regexp.MustCompile(fmt.Sprintf(`(%s)`, PodName)).FindStringSubmatch(content)

	if matchHTTPCode && matchNamespace && matchPodName != nil {
		log.Printf("Match log - Content: %s", content)

		match := regexp.MustCompile(`([A-Z]+) (/\S*) HTTP`).FindStringSubmatch(content)
		if match == nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, "Cannot retrieve information")
			return
		}
		path := match[len(match)-1]
		method := match[len(match)-2]
		podName := matchPodName[len(matchPodName)-1]

		notify := &alert.Data{
			Receiver:          "notification_manager",
			Status:            "firing",
			Alerts:            alert.Alerts{},
			GroupLabels:       alert.KV{"alertname": AlertName, "namespace": Namespace},
			CommonLabels:      alert.KV{"alertname": AlertName, "namespace": Namespace, "severity": Severity},
			CommonAnnotations: alert.KV{},
			ExternalURL:       "",
		}
		alt := alert.Alert{
			Status: "firing",
			Labels: alert.KV{
				"alertname": AlertName,
				"namespace": Namespace,
				"severity":  Severity,
				"pod":       podName,
				"path":      path,
				"method":    method,
			},
			Annotations:  alert.KV{},
			StartsAt:     time.Now(),
			EndsAt:       time.Time{},
			GeneratorURL: "",
			Fingerprint:  "",
		}
		notify.Alerts = append(notify.Alerts, alt)
		notifyBytes, _ := json.Marshal(notify)

		req, err := http.NewRequest("POST", NotificationManagerAddress, bytes.NewReader(notifyBytes))
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, "Failed to create request")
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, "Failed to send request to notification manager")
			return
		}

		log.Printf("Send log to notification manager")
	}
}
