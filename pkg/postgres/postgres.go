package postgres

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"log"
	"task-microservices/config"
	"time"

	"github.com/Masterminds/squirrel"
)

// TODO - before deploy do new path to root.crt
const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
	_ca                  = "/home/romich-v2/.postgresql/root.crt"
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Conn    *pgx.Conn
}

func New(cfg *config.Config) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(_ca)
	if err != nil {
		panic(err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		panic("Failed to append PEM.")
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	connConfig, err := pgx.ParseConfig(cfg.PostgresUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	connConfig.TLSConfig = &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true,
	}

	for pg.connAttempts > 0 {
		pg.Conn, err = pgx.ConnectConfig(context.Background(), connConfig)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close(ctx context.Context) {
	if p.Conn != nil {
		err := p.Conn.Close(ctx)
		if err != nil {
			return
		}
	}
}
