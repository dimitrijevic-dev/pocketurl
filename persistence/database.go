package persistence

import (
	"context"
	"log"
	"pocketurl/config"
	"pocketurl/router"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)
var conn *pgxpool.Pool
func Connect() {
	var err error
	conn,err = pgxpool.New(context.Background(), config.GetEnv("SUPABASE_CONNECTION"))
	if err != nil { log.Fatalf("Supabase connection failed! %v",err)}
}

func GetDestinationByOrigin(origin string) string {
	var link router.Link
	query := "SELECT * FROM links WHERE origin_url = $1"
	pgxscan.Get(context.Background(),conn,&link,query,origin)
	return link.DestinationUrl
}