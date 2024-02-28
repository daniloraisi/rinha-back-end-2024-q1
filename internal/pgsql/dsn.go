package pgsql

import (
	"bytes"
	"net/url"
	"strconv"
)

type Config struct {
	User    string
	Passwd  string
	Addr    string
	Port    uint64
	DBName  string
	SSLMode string
}

func (cfg *Config) FormatDSN() string {
	var buf bytes.Buffer

	buf.WriteString("postgresql://")

	if len(cfg.User) > 0 {
		buf.WriteString(cfg.User)

		if len(cfg.Passwd) > 0 {
			buf.WriteByte(':')
			buf.WriteString(cfg.Passwd)
		}
		buf.WriteByte('@')
	}

	if len(cfg.Addr) > 0 {
		buf.WriteString(cfg.Addr)

		if cfg.Port > 0 {
			buf.WriteByte(':')
			buf.WriteString(strconv.FormatUint(cfg.Port, 10))
		}
	}

	buf.WriteByte('/')
	buf.WriteString(url.PathEscape(cfg.DBName))

	hasParam := false

	if cfg.SSLMode != "" {
		writeDSNParam(&buf, &hasParam, "sslmode", cfg.SSLMode)
	}

	return buf.String()
}

func writeDSNParam(buf *bytes.Buffer, hasParam *bool, name, value string) {
	buf.Grow(1 + len(name) + 1 + len(value))

	if !*hasParam {
		*hasParam = true
		buf.WriteByte('?')
	} else {
		buf.WriteByte('&')
	}

	buf.WriteString(name)
	buf.WriteByte('=')
	buf.WriteString(value)
}
