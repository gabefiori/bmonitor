package database

import (
	"database/sql"
)

type PageMetrics struct {
	ID             int    `json:"id"`
	URL            string `json:"url"`
	AccessCount    int    `json:"access_count"`
	LastAccessedAt string `json:"last_accessed_at"`
}

func InsertMetric(db *sql.DB, url string) error {
	upsertSQL := `
		INSERT INTO page_metrics (url, access_count)
		VALUES (?, 1)
		ON CONFLICT(url) DO UPDATE SET
			access_count = access_count + 1,
			last_accessed_at = excluded.last_accessed_at;`

	_, err := db.Exec(upsertSQL, url)

	if err != nil {
		return err
	}

	return nil
}

func RetriveMetrics(db *sql.DB) ([]*PageMetrics, error) {
	query := `SELECT id, url, access_count, last_accessed_at FROM page_metrics`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var metrics []*PageMetrics

	for rows.Next() {
		var metric PageMetrics

		err := rows.Scan(&metric.ID, &metric.URL, &metric.AccessCount, &metric.LastAccessedAt)

		if err != nil {
			return nil, err
		}

		metrics = append(metrics, &metric)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return metrics, nil
}
