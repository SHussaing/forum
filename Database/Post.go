package Database

type Post struct {
	ID         int
	Title      string
	Content    string
	Categories []Category
}

type Category struct {
	ID   int
	Name string
}

func GetAllPosts() ([]Post, error) {
	postRows, err := Db.Query("SELECT post_ID, title, content FROM Post ORDER BY post_ID DESC")
	if err != nil {
		return nil, err
	}
	defer postRows.Close()

	var posts []Post
	for postRows.Next() {
		var post Post
		err := postRows.Scan(&post.ID, &post.Title, &post.Content)
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
