package logs_handler_function

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
	alert "github.com/prometheus/alertmanager/template"
)

const (
	HTTPCodeNotFound = "404"
	Namespace        = "demo-project"
	PodName          = "wordpress-v1-[A-Za-z0-9]{5,15}-[A-Za-z0-9]{3,10}"
	AlertName        = "404 Request"
	Severity         = "warning"
)

func LogsHandler(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	content := string(in)
	matchHTTPCode, _ := regexp.MatchString(fmt.Sprintf(" %s ", HTTPCodeNotFound), content)
	matchNamespace, _ := regexp.MatchString(fmt.Sprintf("namespace_name\":\"%s", Namespace), content)
	matchPodName := regexp.MustCompile(fmt.Sprintf(`(%s)`, PodName)).FindStringSubmatch(content)

	if matchHTTPCode && matchNamespace && matchPodName != nil {
		log.Printf("Match log - Content: %s", content)

		match := regexp.MustCompile(`([A-Z]+) (/\S*) HTTP`).FindStringSubmatch(content)
		if match == nil {
			return ctx.ReturnOnInternalError(), errors.New("failed to match event")
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

		if _, err := ctx.Send("notify", notifyBytes); err != nil {
			panic(err)
		}
		log.Printf("Send log to notification manager.")
	}
	return ctx.ReturnOnSuccess(), nil
}
