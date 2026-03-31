package repositories

import (
    "database/sql"
    "TweetService/models"
)

type TagRepository interface {
    Upsert(name string) (int, error)
    AssociateTweetTag(tweetID, tagID int) error
    GetAllTags(limit, offset int) ([]*models.Tag, error)
    GetTagWithTweets(name string) ([]*models.Tweet, error)
    GetTopTags(limit int) ([]*models.Tag, error)
    DeleteTag(id int) error
}

type TagRepositoryImpl struct {
    db *sql.DB
}

func NewTagRepository(_db *sql.DB) TagRepository {
    return &TagRepositoryImpl{db: _db}
}

func (r *TagRepositoryImpl) Upsert(name string) (int, error) {
    res, err := r.db.Exec("INSERT INTO tags (name) VALUES (?) ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)", name)
    if err != nil {
        return 0, err
    }
     id, err := res.LastInsertId()
    if err != nil {
        return 0, err
    }
    
    return int(id), nil
}

func (r *TagRepositoryImpl) AssociateTweetTag(tweetID, tagID int) error {
    _, err := r.db.Exec("INSERT IGNORE INTO tweet_tags (tweet_id, tag_id) VALUES (?, ?)", tweetID, tagID)
    return err
}

func (r *TagRepositoryImpl) GetAllTags(limit, offset int) ([]*models.Tag, error) {
    rows, err := r.db.Query("SELECT id, name FROM tags ORDER BY name ASC LIMIT ? OFFSET ?", limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tags []*models.Tag
    for rows.Next() {
        t := &models.Tag{}
        rows.Scan(&t.ID, &t.Name)
        tags = append(tags, t)
    }
    return tags, nil
}

func (r *TagRepositoryImpl) GetTagWithTweets(name string) ([]*models.Tweet, error) {
    query := `
        SELECT t.id, t.user_id, t.tweet, t.createdAt, t.updatedAt
        FROM tweets t
        JOIN tweet_tags tt ON t.id = tt.tweet_id
        JOIN tags h ON h.id = tt.tag_id
        WHERE h.name = ?
    `
    rows, err := r.db.Query(query, name)
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
    return tweets, nil
}

func (r *TagRepositoryImpl) GetTopTags(limit int) ([]*models.Tag, error) {
    query := `
        SELECT h.id, h.name, COUNT(tt.tweet_id) as usage_count
        FROM tags h
        JOIN tweet_tags tt ON h.id = tt.tag_id
        GROUP BY h.id
        ORDER BY usage_count DESC
        LIMIT ?
    `
    rows, err := r.db.Query(query, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tags []*models.Tag
    for rows.Next() {
        t := &models.Tag{}
        rows.Scan(&t.ID, &t.Name, &t.Count)
        tags = append(tags, t)
    }
    return tags, nil
}

func (r *TagRepositoryImpl) DeleteTag(id int) error {
    _, err := r.db.Exec("DELETE FROM tags WHERE id = ?", id)
    return err
}
