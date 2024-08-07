CREATE TABLE IF NOT EXISTS User (
    user_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Category (
    category_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Post (
    post_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    user_ID INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL CHECK (LENGTH(content) <= 5000),
    image BLOB, 
    FOREIGN KEY (user_ID) REFERENCES User(user_ID)
);


CREATE TABLE IF NOT EXISTS Comment (
    comment_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    post_ID INTEGER NOT NULL,
    user_ID INTEGER NOT NULL,
    content TEXT NOT NULL CHECK (LENGTH(content) <= 2000),
    FOREIGN KEY (post_ID) REFERENCES Post(post_ID),
    FOREIGN KEY (user_ID) REFERENCES User(user_ID)
);

CREATE TABLE IF NOT EXISTS Post_Likes (
    user_ID INTEGER,
    post_ID INTEGER,
    status VARCHAR(255) NOT NULL CHECK (status IN ('like', 'dislike')),
    PRIMARY KEY (user_ID, post_ID),
    FOREIGN KEY (user_ID) REFERENCES User(user_ID),
    FOREIGN KEY (post_ID) REFERENCES Post(post_ID)
);

CREATE TABLE IF NOT EXISTS Comment_Likes (
    user_ID INTEGER,
    comment_ID INTEGER,
    status VARCHAR(255) NOT NULL CHECK (status IN ('like', 'dislike')),
    PRIMARY KEY (user_ID, comment_ID),
    FOREIGN KEY (user_ID) REFERENCES User(user_ID),
    FOREIGN KEY (comment_ID) REFERENCES Comment(comment_ID)
);

CREATE TABLE IF NOT EXISTS Post_Categories (
    post_ID INTEGER NOT NULL,
    category_ID INTEGER NOT NULL,
    PRIMARY KEY (post_ID, category_ID),
    FOREIGN KEY (post_ID) REFERENCES Post(post_ID),
    FOREIGN KEY (category_ID) REFERENCES Category(category_ID)
);

CREATE TABLE IF NOT EXISTS Session (
    session_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_ID INTEGER NOT NULL,
    token TEXT NOT NULL,
    expires_at DATETIME NOT NULL,
    FOREIGN KEY (user_ID) REFERENCES User(user_ID)
);

