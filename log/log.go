package log

import (
	"context"
	"fmt"
	"os"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/namo-io/go-kit/constant"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

var globalConfig *Configuration

// ElasticSearchConfiguration elastic search configuration for logging hook
type ElasticSearchConfiguration struct {
	Enabled   bool
	URL       string
	Sniff     bool
	IndexName string
}

// Configuration logger configuration
type Configuration struct {
	// Verbose print include debug messages
	Verbose bool

	// DisableColors print without colors
	DisableColors bool

	// ElasticSearchURL add log to elastic search
	ElasticSearchConfig *ElasticSearchConfiguration
}

// Logger logger interface
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})

	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})

	With(string, interface{}) Logger
	WithContext(context.Context) Logger
	WithError(error) Logger
}

func init() {
	globalConfig = defaultConfiguration()
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		Verbose:             true,
		DisableColors:       false,
		ElasticSearchConfig: nil,
	}
}

// SetGlobalConfiguration set global configuration vriable
func SetGlobalConfiguration(cfg *Configuration) {
	globalConfig = cfg
}

// baseLogger is logrus (github.com/sirupsen/logrus)
type baseLogger struct {
	*logrus.Entry
}

// With returns new logger with custom field
func (l *baseLogger) With(key string, value interface{}) Logger {
	return &baseLogger{l.WithField(key, value)}
}

// With returns new logger with custom field
func (l *baseLogger) WithContext(ctx context.Context) Logger {
	logger := l.Entry

	if app := ctx.Value(constant.ContextApplication); app != nil {
		logger = logger.WithField("Application", app)
	}

	if reqID := ctx.Value(constant.ContextRequestID); reqID != nil {
		logger = logger.WithField("RequestID", reqID)
	}

	return &baseLogger{logger}
}

// WithError returns a new logger with error data set
func (l *baseLogger) WithError(err error) Logger {
	return &baseLogger{l.Entry.WithError(err)}
}

// New create logger without configuration
func New() Logger {
	if globalConfig == nil {
		globalConfig = defaultConfiguration()
	}

	return NewWithConfig(globalConfig)
}

// NewWithConfig create logger with configuration
func NewWithConfig(cfg *Configuration) Logger {
	log := logrus.New()

	// set verbose default: info
	log.SetLevel(logrus.InfoLevel)
	if cfg.Verbose {
		log.SetLevel(logrus.TraceLevel)
	}

	// set method name visible
	log.SetReportCaller(true)

	// set formatter
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		NoColors:        cfg.DisableColors,
		FieldsOrder:     []string{"Hostname", "Application", "RequestPath", "RequestID"},
		TimestampFormat: "2006-01-02 15:04:05.000",
		CustomCallerFormatter: func(f *runtime.Frame) string {
			return fmt.Sprintf("\t %s", f.File)
		},
	})

	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}

	// regist elasticsearch hook
	if cfg.ElasticSearchConfig != nil && cfg.ElasticSearchConfig.Enabled {
		client, err := elastic.NewClient(
			elastic.SetURL(cfg.ElasticSearchConfig.URL),
			elastic.SetSniff(cfg.ElasticSearchConfig.Sniff))
		if err != nil {
			log.Error(err)
		}
		hook, err := elogrus.NewAsyncElasticHook(client, hostname, logrus.DebugLevel, cfg.ElasticSearchConfig.IndexName)
		if err != nil {
			log.Error(err)
		}

		log.Hooks.Add(hook)
	}

	// set application name and hostname
	logger := log.WithFields(map[string]interface{}{
		"Hostname": hostname,
	})

	return &baseLogger{logger}
}
