-- User Table
INSERT INTO User (email, username, password) VALUES 
('john.doe@example.com', 'johndoe', 'password123'),
('jane.smith@example.com', 'janesmith', 'password456'),
('alice.jones@example.com', 'alicejones', 'password789'),
('bob.brown@example.com', 'bobbrown', 'password321'),
('eve.white@example.com', 'evewhite', 'password654'),
('mike.green@example.com', 'mikegreen', 'password987'),
('sara.black@example.com', 'sarablack', 'password321'),
('chris.blue@example.com', 'chrisblue', 'password654');

-- Category Table
INSERT INTO Category (name) VALUES 
('Technology'),
('Science'),
('Arts'),
('Sports'),
('Education'),
('Health'),
('Travel'),
('Food');

-- Post Table
INSERT INTO Post (user_ID, title, content) VALUES 
(1, 'The Future of AI', 'Artificial Intelligence (AI) is rapidly evolving...'),
(2, 'Space Exploration', 'The exploration of space is a field that continues to expand...'),
(3, 'The Beauty of Painting', 'Painting is a form of expression that has been around for centuries...'),
(4, 'The Importance of Physical Fitness', 'Maintaining physical fitness is crucial for a healthy lifestyle...'),
(5, 'Innovations in Education', 'Education has seen numerous innovations in recent years...'),
(6, 'Healthy Eating Habits', 'Healthy eating is essential for maintaining good health...'),
(7, 'Top Travel Destinations', 'Exploring new places can be an enriching experience...'),
(8, 'Delicious Recipes to Try', 'Cooking can be a fun and rewarding hobby...');

-- Comment Table
INSERT INTO Comment (post_ID, user_ID, content) VALUES 
(1, 2, 'Great insights on AI!'),
(2, 1, 'I am excited about space exploration too!'),
(3, 4, 'Painting is indeed a beautiful form of art.'),
(4, 3, 'Physical fitness is so important for everyone.'),
(5, 5, 'Education needs constant innovation.'),
(6, 6, 'Healthy eating is key to a long life.'),
(7, 7, 'Traveling opens up new horizons.'),
(8, 8, 'These recipes look delicious!'),
(1, 3, 'AI will change the world.'),
(2, 4, 'Space is the final frontier.'),
(3, 5, 'Art is a reflection of culture.'),
(4, 6, 'Fitness improves overall well-being.'),
(5, 7, 'Education shapes the future.'),
(6, 8, 'Good nutrition is essential.'),
(7, 1, 'Travel broadens the mind.'),
(8, 2, 'Cooking is a great skill to have.');

-- Post_Likes Table
INSERT INTO Post_Likes (user_ID, post_ID, status) VALUES 
(1, 1, 'like'),
(2, 1, 'like'),
(3, 2, 'dislike'),
(4, 3, 'like'),
(5, 4, 'like'),
(6, 5, 'dislike'),
(7, 6, 'like'),
(8, 7, 'like'),
(1, 8, 'like'),
(2, 2, 'like'),
(3, 3, 'dislike'),
(4, 4, 'like'),
(5, 5, 'like'),
(6, 6, 'dislike'),
(7, 7, 'like'),
(8, 8, 'like');

-- Comment_Likes Table
INSERT INTO Comment_Likes (user_ID, comment_ID, status) VALUES 
(1, 1, 'like'),
(2, 2, 'like'),
(3, 3, 'dislike'),
(4, 4, 'like'),
(5, 5, 'like'),
(6, 6, 'dislike'),
(7, 7, 'like'),
(8, 8, 'like'),
(1, 9, 'like'),
(2, 10, 'like'),
(3, 11, 'dislike'),
(4, 12, 'like'),
(5, 13, 'like'),
(6, 14, 'dislike'),
(7, 15, 'like'),
(8, 16, 'like');

-- Post_Categories Table
INSERT INTO Post_Categories (post_ID, category_ID) VALUES 
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(1, 2),
(2, 3),
(3, 4),
(4, 5),
(5, 6),
(6, 7),
(7, 8),
(8, 1);
