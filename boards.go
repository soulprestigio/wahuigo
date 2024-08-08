package boards

import (
 "fmt"
 "database/sql"
 "html/template"
 _ "github.com/mattn/go-sqlite3"
)

// Board represents a board in the database
type Board struct {
 ID          int    `json:"id"`
 Name        string `json:"name"`
 Description string `json:"description"`
 URL         string `json:"url"`
}

// GetBoards retrieves boards from the database and renders them using a template.
func GetAllBoards(db *sql.DB, tmpl *template.Template) ([]Board, error) { 
 rows, err := db.Query("SELECT * FROM boards")
 if err != nil {
  return nil, fmt.Errorf("failed to query boards: %w", err) // Use fmt.Errorf for better error messages
 }
 defer rows.Close()

 var boards []Board
 for rows.Next() {
  var board Board
  err := rows.Scan(&board.ID, &board.Name, &board.Description, &board.URL)
  if err != nil {
   return nil, fmt.Errorf("failed to scan board: %w", err)
  }
  boards = append(boards, board)
 }
 if err := rows.Err(); err != nil {
  return nil, fmt.Errorf("error iterating over boards: %w", err)
 }

 return boards, nil
}

// AddBoard adds a new board to the database.
func AddBoard(db *sql.DB, name, description, url string) error {
 stmt, err := db.Prepare("INSERT INTO boards(name, description, url) VALUES(?, ?, ?)")
 if err != nil {
  return fmt.Errorf("failed to prepare statement: %w", err)
 }
 defer stmt.Close()

 _, err = stmt.Exec(name, description, url)
 if err != nil {
  return fmt.Errorf("failed to execute statement: %w", err)
 }

 return nil
}
