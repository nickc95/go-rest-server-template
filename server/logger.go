package server

import (
	logrus "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"

	"go-rest-server-template/models"
)

var Log *logrus.Logger

func loggerInit() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "type",
		},
	})
	Log.SetOutput(&lumberjack.Logger{
		Filename: "logs/" + configObj.ApplicationName + ".log",
		MaxSize:  50, // mb
	})
}

// TODO: continue to add more log data, be sure to mask confidential data
func log(w models.ResponseContextWriter, r *http.Request, requestData *models.RequestData) {
	Log.WithFields(logrus.Fields{
		"application": "",
		"code":        w.Status,
		"ip":          r.RemoteAddr,
		"userAgent":   r.UserAgent(),
		"requestPath": r.Host + r.URL.EscapedPath(),
		"clientError": w.ClientError,
	}).Error(w.ErrorMessage)
}
