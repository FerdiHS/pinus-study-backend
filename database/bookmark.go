package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Bookmark a thread
func BookmarkThread(db *sql.DB, threadId int, userId int) error {
	sql_statement := `
	INSERT INTO bookmark_threads(thread_id, user_id)
	VALUES ($1, $2);
	`
	_, err := db.Exec(sql_statement, threadId, userId)
	return err
}

// Unbookmark a thread
func UnbookmarkThread(db *sql.DB, threadId int, userId int) error {
	sql_statement := `
	DELETE FROM bookmark_threads
	WHERE thread_id = $1 AND user_id = $2;
	`
	_, err := db.Exec(sql_statement, threadId, userId)
	return err
}

// Get whether a thread is being bookmarked by the user
func GetBookmarkThread(db *sql.DB, threadId int, userId int) (bool, error) {
	sql_statement := `
	SELECT COUNT(*) > 0 AS is_bookmarked FROM bookmark_threads
	WHERE thread_id = $1 AND user_id = $2;
	`

	var isBookmarked bool
	err := db.QueryRow(sql_statement, threadId, userId).Scan(&isBookmarked)
	return isBookmarked, err
}
