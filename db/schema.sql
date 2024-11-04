CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(20) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    key TEXT,
    ExpireDate DATETIME,
    FOREIGN KEY (username) REFERENCES users(username)

);   

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    content TEXT,
    author VARCHAR(20) NOT NULL,
    category TEXT NOT NULL,
    likesCount INTEGER,
    dislikesCount INTEGER,
    liked BOOLEAN,
    disliked BOOLEAN,
    FOREIGN KEY (author) REFERENCES users(username),
    FOREIGN KEY (category) REFERENCES categories(name)
);


CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY,
    postId INTEGER,
    author VARCHAR(20) NOT NULL,
    content TEXT,
    likesCount INTEGER,
    dislikesCount INTEGER,
    liked BOOLEAN,
    disliked BOOLEAN,
    FOREIGN KEY (postId) REFERENCES posts(id),
    FOREIGN KEY (author) REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS engagement (
    id INTEGER PRIMARY KEY,
    postId INTEGER,
    userId INTEGER,
    commentId INTEGER,
    like BOOLEAN,
    dislike BOOLEAN,
    FOREIGN KEY (postId) REFERENCES posts(id),
    FOREIGN KEY (userId) REFERENCES users(id),
    FOREIGN KEY (commentId) REFERENCES comments(id)
);
