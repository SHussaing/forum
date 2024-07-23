package Database

type Category struct {
	ID   int
	Name string
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
