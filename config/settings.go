package config

import (
	"encoding/json"
	"fmt"
)

type environment string

const (
	Dev  environment = "dev"
	Prod environment = "prod"
)

func (e environment) String() string {
	return string(e)
}

func (e *environment) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch t := v.(type) {
	case string:
		env := environment(t)
		if env != Dev && env != Prod {
			return fmt.Errorf("unsupported environment %s. Options are: %s - %s", env, Prod, Dev)
		}
		*e = env
		return nil

	default:
		return fmt.Errorf("environment must be of type string")
	}
}

type level string

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
)

func (l level) String() string {
	return string(l)
}

func (l *level) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch t := v.(type) {
	case string:
		lev := level(t)
		if lev != Debug && lev != Info && lev != Warn && lev != Error {
			return fmt.Errorf("invalid log level %s. Options are: %s - %s - %s - %s", lev, Debug, Info, Warn, Error)
		}
		*l = lev
		return nil

	default:
		return fmt.Errorf("log level must be of type string")
	}
}
