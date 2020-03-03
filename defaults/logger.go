package defaults

import (
	"github.com/sirupsen/logrus"
)

type Logrus struct {
	*logrus.Entry
}

func NewLogrus() *Logrus {
	return &Logrus{Entry: logrus.NewEntry(logrus.New())}
}

func (l *Logrus) logBuilder(url string, params map[string]interface{}, err error) *logrus.Entry {
	log := l.Entry
	if err != nil {
		log.WithError(err)
	}
	if params != nil {
		log.WithFields(params)
	}
	if url != "" {
		log.WithField("url", url)
	}
	return log
}

func (l *Logrus) Panic(msg string, url string, params map[string]interface{}, err error) {
	l.logBuilder(url, params, err).Panic(msg)
}

func (l *Logrus) Error(msg string, url string, params map[string]interface{}, err error) {
	l.logBuilder(url, params, err).Error(msg)
}

func (l *Logrus) Warn(msg string, url string, params map[string]interface{}, err error) {
	l.logBuilder(url, params, err).Error(msg)
}

func (l *Logrus) Info(msg string, url string, params map[string]interface{}, err error) {
	l.logBuilder(url, params, err).Info(msg)
}

func (l *Logrus) Debug(msg string, url string, params map[string]interface{}, err error) {
	l.logBuilder(url, params, err).Debug(msg)
}

func (l *Logrus) Trace(msg string, url string, params map[string]interface{}, err error) {
	l.logBuilder(url, params, err).Trace(msg)
}
