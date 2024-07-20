CREATE TABLE User (
    user_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE Category (
    category_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE Post (
    post_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    user_ID INTEGER NOT NULL,
    category_ID INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_ID) REFERENCES User(user_ID),
    FOREIGN KEY (category_ID) REFERENCES Category(category_ID)
);

CREATE TABLE Comment (
    comment_ID INTEGER PRIMARY KEY AUTOINCREMENT,
    post_ID INTEGER NOT NULL,
    user_ID INTEGER NOT NULL,
    content VARCHAR(255) NOT NULL,
    FOREIGN KEY (post_ID) REFERENCES Post(post_ID),
    FOREIGN KEY (user_ID) REFERENCES User(user_ID)
);

CREATE TABLE Post_Likes (
    user_ID INTEGER,
    post_ID INTEGER,
    status VARCHAR(255),
    PRIMARY KEY (user_ID, post_ID),
    FOREIGN KEY (user_ID) REFERENCES User(user_ID),
    FOREIGN KEY (post_ID) REFERENCES Post(post_ID)
);

CREATE TABLE Comment_Likes (
    user_ID INTEGER,
    comment_ID INTEGER,
    status VARCHAR(255),
    PRIMARY KEY (user_ID, comment_ID),
    FOREIGN KEY (user_ID) REFERENCES User(user_ID),
    FOREIGN KEY (comment_ID) REFERENCES Comment(comment_ID)
);
