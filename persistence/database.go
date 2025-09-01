package persistence

import (
	"context"
	"log"
	"pocketurl/config"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Link struct {
	ID int8 `json:"id"`
	OriginUrl string `json:"origin_url"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	DestinationUrl string `json:"destination_url"`
}


var conn *pgxpool.Pool
func Start() {
	go startCleanupScheduler() 

	var err error
	conn,err = pgxpool.New(context.Background(), config.GetEnv("SUPABASE_CONNECTION"))
	if err != nil { log.Fatalf("Supabase connection failed! %v",err)}
}

func GetLinkByOrigin(origin string) *Link {
	var link Link
	query := `SELECT * FROM links WHERE origin_url = $1`
	pgxscan.Get(context.Background(),conn,&link,query,origin)
	return &link
}

func AddLink(link Link) error {
	query := `INSERT INTO links (origin_url, destination_url, expires_at) VALUES ($1, $2, $3)`
	_,err := conn.Exec(context.Background(),query,link.OriginUrl,link.DestinationUrl,link.ExpiresAt)
	return err
}

func DeleteLink(link Link) error {
	query := `DELETE FROM links WHERE origin_url = $1`
	_,err := conn.Exec(context.Background(),query,link.OriginUrl)
	return err
}

//This is meant to run hourly to purge expired links as well
func CleanExpiredLinks() error {
	query := `DELETE FROM links WHERE expires_at < NOW()`
	_, err := conn.Exec(context.Background(),query)
	return err
}