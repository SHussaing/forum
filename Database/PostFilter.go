package Database

import (
	"strings"
)

func GetFilteredPosts(userID int, filters, categoryIDs []string, isLoggedIn bool) ([]Post, error) {
	query := `
		SELECT p.post_ID, p.user_ID, p.title, p.content, u.username
		FROM Post p
		JOIN User u ON p.user_ID = u.user_ID`
	var conditions []string
	var args []interface{}

	if isLoggedIn && len(filters) > 0 {
		for _, filter := range filters {
			switch filter {
			case "my_posts":
				conditions = append(conditions, "p.user_ID = ?")
				args = append(args, userID)
			case "liked_posts":
				conditions = append(conditions, "p.post_ID IN (SELECT post_ID FROM Post_Likes WHERE user_ID = ? AND status = 'like')")
				args = append(args, userID)
			}
		}
	}

	if len(categoryIDs) > 0 {
		placeholders := make([]string, len(categoryIDs))
		for i, id := range categoryIDs {
			placeholders[i] = "?"
			args = append(args, id)
		}
		conditions = append(conditions, "p.post_ID IN (SELECT post_ID FROM Post_Categories WHERE category_ID IN ("+strings.Join(placeholders, ",")+"))")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY p.post_ID DESC"

	return queryPosts(query, args...)
}

func queryPosts(query string, args ...interface{}) ([]Post, error) {
	rows, err := Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Username)
		if err != nil {
			return nil, err
		}

		categoryRows, err := Db.Query(`
            SELECT c.category_ID, c.name
            FROM Category c
            JOIN Post_Categories pc ON c.category_ID = pc.category_ID
            WHERE pc.post_ID = ?`, post.ID)
		if err != nil {
			return nil, err
		}
		defer categoryRows.Close()

		var categories []Category
		for categoryRows.Next() {
			var category Category
			err := categoryRows.Scan(&category.ID, &category.Name)
			if err != nil {
				return nil, err
			}
			categories = append(categories, category)
		}
		post.Categories = categories

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
