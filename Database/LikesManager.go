package Database

import "database/sql"

// LikePost handles liking a post
func LikePost(userID, postID int) error {
	return handlePostLike(userID, postID, "like")
}

// DislikePost handles disliking a post
func DislikePost(userID, postID int) error {
	return handlePostLike(userID, postID, "dislike")
}

// LikeComment handles liking a comment
func LikeComment(userID, commentID int) error {
	return handleCommentLike(userID, commentID, "like")
}

// DislikeComment handles disliking a comment
func DislikeComment(userID, commentID int) error {
	return handleCommentLike(userID, commentID, "dislike")
}

// handlePostLike is a helper function to handle both like and dislike actions for posts
func handlePostLike(userID, postID int, action string) error {
	var status string
	err := Db.QueryRow("SELECT status FROM Post_Likes WHERE user_ID = ? AND post_ID = ?", userID, postID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			// Insert new like or dislike
			_, err = Db.Exec("INSERT INTO Post_Likes (user_ID, post_ID, status) VALUES (?, ?, ?)", userID, postID, action)
			return err
		}
		return err
	}

	if status == action {
		// Delete the existing like or dislike
		_, err = Db.Exec("DELETE FROM Post_Likes WHERE user_ID = ? AND post_ID = ?", userID, postID)
		return err
	}

	// Update the existing like or dislike
	_, err = Db.Exec("UPDATE Post_Likes SET status = ? WHERE user_ID = ? AND post_ID = ?", action, userID, postID)
	return err
}

// handleCommentLike is a helper function to handle both like and dislike actions for comments
func handleCommentLike(userID, commentID int, action string) error {
	var status string
	err := Db.QueryRow("SELECT status FROM Comment_Likes WHERE user_ID = ? AND comment_ID = ?", userID, commentID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			// Insert new like or dislike
			_, err = Db.Exec("INSERT INTO Comment_Likes (user_ID, comment_ID, status) VALUES (?, ?, ?)", userID, commentID, action)
			return err
		}
		return err
	}

	if status == action {
		// Delete the existing like or dislike
		_, err = Db.Exec("DELETE FROM Comment_Likes WHERE user_ID = ? AND comment_ID = ?", userID, commentID)
		return err
	}

	// Update the existing like or dislike
	_, err = Db.Exec("UPDATE Comment_Likes SET status = ? WHERE user_ID = ? AND comment_ID = ?", action, userID, commentID)
	return err
}

// GetPostLikes returns the number of likes and dislikes for a post
func GetPostLikes(postID int) (int, int, error) {
	var likes, dislikes int
	err := Db.QueryRow("SELECT COUNT(*) FROM Post_Likes WHERE post_ID = ? AND status = 'like'", postID).Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	err = Db.QueryRow("SELECT COUNT(*) FROM Post_Likes WHERE post_ID = ? AND status = 'dislike'", postID).Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}

// GetCommentLikes returns the number of likes and dislikes for a comment
func GetCommentLikes(commentID int) (int, int, error) {
	var likes, dislikes int
	err := Db.QueryRow("SELECT COUNT(*) FROM Comment_Likes WHERE comment_ID = ? AND status = 'like'", commentID).Scan(&likes)
	if err != nil {
		return 0, 0, err
	}

	err = Db.QueryRow("SELECT COUNT(*) FROM Comment_Likes WHERE comment_ID = ? AND status = 'dislike'", commentID).Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}