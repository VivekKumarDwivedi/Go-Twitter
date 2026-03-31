package repositories

import(
	"database/sql"
	"TweetService/models"
    "fmt"
)

type TweetRepository interface {
	Create(tweet models.Tweet) (*models.Tweet, error)
	GetByID(id int) (*models.Tweet, error)
	GetAll(limit, offset int,userID *int,hashTags *string, search *string) ([]models.Tweet, error)
	UpdateTweet(id int,content string) error
	DeleteTweet(id int) error
}

type TweetRepositoryImpl struct {
	db *sql.DB
}


func NewTweetRepository(db *sql.DB) TweetRepository {
	return &TweetRepositoryImpl{
		db: db,
	}
}


func (r *TweetRepositoryImpl) Create(tweet models.Tweet) (*models.Tweet, error) {
	query := "INSERT INTO tweets (user_id,tweet) VALUES (?, ?)"

	res, err := r.db.Exec(query,tweet.UserID,tweet.Tweet)

	if err != nil {
		return nil, err
	}

	id, _:= res.LastInsertId()

	tweet.ID = int(id)

	return &tweet,nil
}

func (r *TweetRepositoryImpl) GetByID(id int) (*models.Tweet, error) {
	 if id <= 0 {
        return nil, fmt.Errorf("invalid tweet ID: must be greater than zero")
    }

    query := "SELECT id, user_id, tweet, createdAt, updatedAt FROM tweets WHERE id = ?"
    row := r.db.QueryRow(query, id)

    var tw models.Tweet
    err := row.Scan(&tw.ID, &tw.UserID, &tw.Tweet, &tw.CreatedAt, &tw.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            // No tweet found for this ID
            return nil, fmt.Errorf("no tweet found with id %d", id)
        }
        // Unexpected DB error
        return nil, fmt.Errorf("error fetching tweet by id %d: %w", id, err)
    }

    return &tw, nil
}

func (r *TweetRepositoryImpl) GetAll(limit, offset int, userID *int, hashTags *string, search *string) ([]models.Tweet, error) {
	 query := "SELECT id, user_id, tweet, createdAt, updatedAt FROM tweets WHERE 1=1"
    args := []interface{}{}

    if userID != nil {
        query += " AND user_id = ?"
        args = append(args, *userID)
    }
    if search != nil && *search != "" {
        query += " AND tweet LIKE ?"
        args = append(args, "%"+*search+"%")
    }
    if hashTags != nil && *hashTags != "" {
        query += ` AND id IN (
            SELECT tt.tweet_id FROM tweet_tags tt
            JOIN tags h ON h.id = tt.tag_id
            WHERE h.name = ?
        )`
        args = append(args, *hashTags)
    }

    query += " ORDER BY createdAt DESC LIMIT ? OFFSET ?"
    args = append(args, limit, offset)

    rows, err := r.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tweets []*models.Tweet
    for rows.Next() {
        tw := &models.Tweet{}
        rows.Scan(&tw.ID, &tw.UserID, &tw.Tweet, &tw.CreatedAt, &tw.UpdatedAt)
        tweets = append(tweets, tw)
    }
    result := make([]models.Tweet,len(tweets))
    for i, tweet := range tweets {
        result[i]=*tweet
    }
    return result, nil
}

func (r *TweetRepositoryImpl) UpdateTweet(id int, content string) error {
	_, err := r.db.Exec("UPDATE tweets SET tweet = ?, updatedAt = CURRENT_TIMESTAMP WHERE id = ?",content,id)
	return err
}

func (r *TweetRepositoryImpl) DeleteTweet(id int) error {
	_, err := r.db.Exec("DELETE FROM tweets WHERE id = ?",id)
	return err
}