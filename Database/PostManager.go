package Database

import (
	"database/sql"
	"errors"
)

type Post struct {
	ID         int
	UserID     int
	Title      string
	Content    string
	Categories []Category
	Username   string
	Likes      int
	Dislikes   int
}

type Category struct {
	ID   int
	Name string
}

type Comment struct {
	CommentID int
	PostID    int
	UserID    int
	Content   string
	Username  string
	Likes     int
	Dislikes  int
}

func GetAllPosts() ([]Post, error) {
	query := `
        SELECT p.post_ID, p.user_ID, p.title, p.content, u.username
        FROM Post p
        JOIN User u ON p.user_ID = u.user_ID
        ORDER BY p.post_ID DESC`

	postRows, err := Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer postRows.Close()

	var posts []Post
	for postRows.Next() {
		var post Post
		err := postRows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Username)
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

	if err = postRows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPostAndComments(postID int) (Post, []Comment, error) {
	var post Post
	var comments []Comment

	// Get the post details along with the username of the author, likes, and dislikes
	err := Db.QueryRow(`
		SELECT p.post_ID, p.user_ID, p.title, p.content, u.username,
		       (SELECT COUNT(*) FROM Post_Likes WHERE post_ID = p.post_ID AND status = 'like') AS likes,
		       (SELECT COUNT(*) FROM Post_Likes WHERE post_ID = p.post_ID AND status = 'dislike') AS dislikes
		FROM Post p
		JOIN User u ON p.user_ID = u.user_ID
		WHERE p.post_ID = ?`, postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Username, &post.Likes, &post.Dislikes)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, comments, errors.New("post not found")
		}
		return post, comments, err
	}

	// Get the comments for the post along with likes and dislikes
	rows, err := Db.Query(`
		SELECT c.comment_ID, c.post_ID, c.user_ID, c.content, u.username,
		       (SELECT COUNT(*) FROM Comment_Likes WHERE comment_ID = c.comment_ID AND status = 'like') AS likes,
		       (SELECT COUNT(*) FROM Comment_Likes WHERE comment_ID = c.comment_ID AND status = 'dislike') AS dislikes
		FROM Comment c
		JOIN User u ON c.user_ID = u.user_ID
		WHERE c.post_ID = ?`, postID)
	if err != nil {
		return post, comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &comment.Username, &comment.Likes, &comment.Dislikes)
		if err != nil {
			return post, comments, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return post, comments, err
	}

	return post, comments, nil
}

func GetAllCategories() ([]Category, error) {
	rows, err := Db.Query("SELECT category_ID, name FROM Category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func CreatePost(title, content string, userID int) (int, error) {
	result, err := Db.Exec("INSERT INTO Post (title, content, user_ID) VALUES (?, ?, ?)", title, content, userID)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(postID), nil
}

func AddPostCategory(postID, categoryID int) error {
	_, err := Db.Exec("INSERT INTO Post_Categories (post_ID, category_ID) VALUES (?, ?)", postID, categoryID)
	return err
}

func AddComment(postID, userID int, content string) (int, error) {
	result, err := Db.Exec("INSERT INTO Comment (post_ID, user_ID, content) VALUES (?, ?, ?)", postID, userID, content)
	if err != nil {
		return 0, err
	}
	commentID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(commentID), nil
}

func GetUsernameByID(userID int) (string, error) {
	var username string
	err := Db.QueryRow("SELECT username FROM User WHERE user_ID = ?", userID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
