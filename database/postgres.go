package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"nicolasjosecastro/go/cqrs/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

func (repo *PostgresRepository) InsertFeed(ctx context.Context, feed *models.Feed) error {
	// Insert the feed into the database
	_, err := repo.db.ExecContext(ctx, "INSERT INTO feeds (id, title, description) VALUES ($1, $2, $3)", feed.Title, feed.Title, feed.Description)
	return err
}

func (repo *PostgresRepository) ListFeeds(ctx context.Context) ([]*models.Feed, error) {
	// Query the database for all feeds
	rows, err := repo.db.QueryContext(ctx, "SELECT id, title, description, create_at FROM feeds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and create a feed for each
	feeds := []*models.Feed{} //Sliece of pointers to Feed
	for rows.Next() {
		feed := &models.Feed{}
		err := rows.Scan(&feed.ID, &feed.Title, &feed.Description, &feed.CreatedAt)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, feed)
	}

	return feeds, nil
}
