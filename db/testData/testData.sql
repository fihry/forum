PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(20) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL
);
INSERT INTO users VALUES(1,'user1','$2a$10$MdxAX0XjvEh0.wUwtgmTBeKzKBlx8.lZJGYUSGugcDD.91mrXjape','user1@example.com');
INSERT INTO users VALUES(2,'user2','$2a$10$cc.RbGxHPmfTIuCYN37HluVsSTNwkAxHLKq6A/qGTvofbH5iiWbgi','user2@example.com');
INSERT INTO users VALUES(3,'user3','$2a$10$yIkDib0of8B4jFzZuG.uSuigXTBm30JB5DH4fjR3vxdgIWVgLjt02','user3@example.com');
INSERT INTO users VALUES(4,'user4','$2a$10$7T.ERBRFGbyoqhfBWQcfZu7JvpQ/m3VBP7MB9hlKcjf1l2BCS/8Ei','user4@example.com');
INSERT INTO users VALUES(5,'user5','$2a$10$dARTGIb7qWykUixoSlSHqOaDNiF2s4hy9Ed7YqewqoM6oXc2eKZoK','user5@example.com');
INSERT INTO users VALUES(6,'user6','$2a$10$IjRjMNB8hT4keiBYqa923.5T1lfGhI1t3.7ngJXtsgqwbtbWwMokC','user6@example.com');
INSERT INTO users VALUES(7,'user7','$2a$10$gu7O.qL1ZODxgf4C8vTp5.0pkGAYQRNrv/Z5hsGAIyHigaRzhkABC','user7@example.com');
INSERT INTO users VALUES(8,'user8','$2a$10$NwzVMnuEJzZIpjP7orQcWuUJIbYlHV7h2FPptpARbYqqxIv3VHHHq','user8@example.com');
INSERT INTO users VALUES(9,'user9','$2a$10$CKcRYraa4mjqPvF1YsiyReo/WkSmuby0yr7EtrWp6HttQtryAIvqW','user9@example.com');
INSERT INTO users VALUES(10,'user10','$2a$10$wOIaMrHQNoEwNuC.zeKnZefxSwPSEKmSzyMni8sIp8iBxB/LnJkWa','user10@example.com');
INSERT INTO users VALUES(11,'user11','$2a$10$iEVZRCzyQ.2j0RlRLTPdj.1lj52OQnEOFTXZ/fYzaHcmvwP7EbmIi','user11@example.com');
INSERT INTO users VALUES(12,'user12','$2a$10$Y6Odw8LNnxCA5maBIo6bfO0S4GkFjCKgo6hkPh2D02okZoqeCKf26','user12@example.com');
INSERT INTO users VALUES(13,'user13','$2a$10$hiwtreocUR7Y0Zq5CJJX..CHKL7Kf5SkKftZ3/QP.sm42prD8KqL2','user13@example.com');
INSERT INTO users VALUES(14,'user14','$2a$10$UD/0KQMdHCk6rP1pL4Ut3OuB8f84cCX9cFGYwQA6uMCGxTOkTLGYS','user14@example.com');
INSERT INTO users VALUES(15,'user15','$2a$10$FeRBIGsj/UuPhba5Bbe.q.RQzV0aOr25S/JLBYc/rYoavTWepilnW','user15@example.com');
INSERT INTO users VALUES(16,'user16','$2a$10$zvEXXLVijoDn1Xw7EsiFkucQ9.qguIRaJouGmY8gT.IL6CIl5PPDW','user16@example.com');
INSERT INTO users VALUES(17,'user17','$2a$10$Q0i.uPvBORSt9occJAD6HeizdZoSRxJ5.Yu04U1gFKG0KzC5YGqG.','user17@example.com');
INSERT INTO users VALUES(18,'user18','$2a$10$VXDgxXp2XKJmssE82VVMNueQUT4i/sMk2qcoANHrnHpdkOH4JD9KS','user18@example.com');
INSERT INTO users VALUES(19,'user19','$2a$10$7NoomaSyINA5cAtPuMACK.HjVYP.HlsNTIQQKKbJ/9t.CuQVtxuka','user19@example.com');
INSERT INTO users VALUES(20,'fihry','$2a$10$xVNf/8uMfPH2d5im4ugHaux.mTXrECWV3sFdiQ/A804qx4bK/T9EO','fihrylar@gmail.com');
CREATE TABLE sessions (
    id INTEGER PRIMARY KEY,
    userId INTEGER NOT NULL,
    key TEXT,
    ExpireDate DATETIME,
    FOREIGN KEY (userId) REFERENCES users(id)

);
INSERT INTO sessions VALUES(1,1,'938b5a82-a1d0-44a3-9633-2adc7636c286','2024-11-03 13:16:57.206495385+01:00');
INSERT INTO sessions VALUES(2,2,'e1b504f1-9c86-4395-a60c-7302c696ee39','2024-11-03 13:17:07.78609285+01:00');
INSERT INTO sessions VALUES(3,3,'9adb7d37-f3d3-4f41-9753-74a43fe73cb6','2024-11-03 13:17:16.908018494+01:00');
INSERT INTO sessions VALUES(4,4,'beb00537-e20d-41c5-84f5-077c47e01e49','2024-11-03 13:19:26.232408505+01:00');
INSERT INTO sessions VALUES(5,5,'fb421bfa-6a0e-434a-a919-ceff00d36342','2024-11-03 13:19:39.892784387+01:00');
INSERT INTO sessions VALUES(6,6,'4b715f9e-608e-42de-bb91-dfdbfad54da2','2024-11-03 13:19:50.074492538+01:00');
INSERT INTO sessions VALUES(7,7,'a65a1d30-89a8-47d8-8aca-d1e5ac51d774','2024-11-03 13:21:13.549909393+01:00');
INSERT INTO sessions VALUES(8,8,'00cdeadf-bf43-4d85-b2da-5c710990db7e','2024-11-03 13:21:22.479023276+01:00');
INSERT INTO sessions VALUES(9,9,'e3d62aba-4206-4eec-af41-7c4737e1fc2e','2024-11-03 13:21:32.490104539+01:00');
INSERT INTO sessions VALUES(10,10,'19cf0ba7-722e-49e1-bb4e-044423aaeee1','2024-11-03 13:21:41.46309955+01:00');
INSERT INTO sessions VALUES(11,11,'c5ae5b76-e405-4c3c-8dd8-21af3f451ec0','2024-11-03 13:21:50.597277823+01:00');
INSERT INTO sessions VALUES(12,12,'f307a429-65bc-4b02-bb4c-ba999d6c700f','2024-11-03 13:22:19.656288902+01:00');
INSERT INTO sessions VALUES(13,13,'d5b1ecb8-4209-4e09-8fd9-f7e2e4937284','2024-11-03 13:22:27.597119068+01:00');
INSERT INTO sessions VALUES(14,14,'87466c9d-b8cd-4109-bdac-21453fd77b84','2024-11-03 13:22:37.101957874+01:00');
INSERT INTO sessions VALUES(15,15,'5dc45630-e174-412f-8d2f-d8ae54e9b3c7','2024-11-03 13:22:49.997017654+01:00');
INSERT INTO sessions VALUES(16,16,'84ef4856-259e-4fb9-adc8-062b0a9f8986','2024-11-03 13:23:01.900068002+01:00');
INSERT INTO sessions VALUES(17,17,'19088e4c-ac7a-480b-9191-76577d6c7bb0','2024-11-03 13:23:12.623416021+01:00');
INSERT INTO sessions VALUES(18,18,'5d3a07fe-5d41-4dee-80ef-39fb815fd746','2024-11-03 13:23:23.146020449+01:00');
INSERT INTO sessions VALUES(19,19,'2c128ac4-cbb1-4d39-8aa3-0a58d5647b15','2024-11-03 13:23:34.037470511+01:00');
INSERT INTO sessions VALUES(21,20,'8619cbeb-0c83-4f26-9952-00fb65749c67','2024-11-03 13:53:26.048970077+01:00');
INSERT INTO sessions VALUES(22,20,'3800b70e-e1f1-4250-ba3b-90cd87d8b9c7','2024-11-03 13:53:34.212430951+01:00');
INSERT INTO sessions VALUES(23,20,'85a94ae7-4b09-484c-9025-bedba0eeb7f8','2024-11-03 14:51:46.815226074+01:00');
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    createdAt DATETIME,
    content TEXT,
    author VARCHAR(20) NOT NULL,
    category INTEGER NOT NULL,
    likesCount INTEGER,
    dislikesCount INTEGER,
    liked BOOLEAN,
    disliked BOOLEAN,
    FOREIGN KEY (author) REFERENCES users(username),
    FOREIGN KEY (category) REFERENCES categories(name)
);
-- Inserting into the posts table with createdAt values
INSERT INTO posts VALUES(1, 'First Post', '2024-11-05 10:00:00', 'This is the content for the first post.', 'user1', 'General', 10, 2, 1, 0);
INSERT INTO posts VALUES(2, 'Second Post', '2024-11-05 10:30:00', 'Content of the second post goes here.', 'user2', 'News', 5, 1, 1, 0);
INSERT INTO posts VALUES(3, 'Third Post', '2024-11-05 11:00:00', 'Exploring the details of the third post.', 'user3', 'Tech', 15, 0, 0, 1);
INSERT INTO posts VALUES(4, 'Fourth Post', '2024-11-05 11:30:00', 'Thoughts on various subjects in the fourth post.', 'user4', 'Lifestyle', 20, 4, 1, 0);
INSERT INTO posts VALUES(5, 'Fifth Post', '2024-11-05 12:00:00', 'Fifth post content discussing important topics.', 'user5', 'Education', 8, 1, 0, 1);
INSERT INTO posts VALUES(6, 'Sixth Post', '2024-11-05 12:30:00', 'Insights shared in the sixth post.', 'user6', 'Travel', 12, 3, 1, 0);
INSERT INTO posts VALUES(7, 'Seventh Post', '2024-11-05 13:00:00', 'Seventh post goes deep into analysis.', 'user7', 'Health', 6, 0, 1, 0);
INSERT INTO posts VALUES(8, 'Eighth Post', '2024-11-05 13:30:00', 'Eighth post discusses current trends.', 'user8', 'Entertainment', 22, 5, 0, 1);
INSERT INTO posts VALUES(9, 'Ninth Post', '2024-11-05 14:00:00', 'A short note on the ninth post.', 'user9', 'General', 3, 2, 1, 0);
INSERT INTO posts VALUES(10, 'Tenth Post', '2024-11-05 14:30:00', 'Content of the tenth post with great insights.', 'user10', 'News', 7, 1, 0, 1);
INSERT INTO posts VALUES(11, 'Eleventh Post', '2024-11-05 15:00:00', 'The eleventh post reflects on key ideas.', 'user11', 'Tech', 18, 2, 1, 0);
INSERT INTO posts VALUES(12, 'Twelfth Post', '2024-11-05 15:30:00', 'Twelfth post content sharing personal views.', 'user12', 'Lifestyle', 5, 3, 0, 1);
INSERT INTO posts VALUES(13, 'Thirteenth Post', '2024-11-05 16:00:00', 'Thoughts in the thirteenth post.', 'user13', 'Education', 11, 0, 1, 0);
INSERT INTO posts VALUES(14, 'Fourteenth Post', '2024-11-05 16:30:00', 'Fourteenth post with exciting travel stories.', 'user14', 'Travel', 14, 1, 0, 1);
INSERT INTO posts VALUES(15, 'Fifteenth Post', '2024-11-05 17:00:00', 'Fifteenth post provides health tips.', 'user15', 'Health', 19, 2, 1, 0);
INSERT INTO posts VALUES(16, 'Sixteenth Post', '2024-11-05 17:30:00', 'A discussion on trends in the sixteenth post.', 'user16', 'Entertainment', 9, 4, 0, 1);
INSERT INTO posts VALUES(17, 'Seventeenth Post', '2024-11-05 18:00:00', 'Seventeenth post addresses common issues.', 'user17', 'General', 13, 1, 1, 0);
INSERT INTO posts VALUES(18, 'Eighteenth Post', '2024-11-05 18:30:00', 'The eighteenth post contains reviews.', 'user18', 'News', 16, 0, 0, 1);
INSERT INTO posts VALUES(19, 'Nineteenth Post', '2024-11-05 19:00:00', 'Insights in the nineteenth post.', 'user19', 'Tech', 12, 5, 1, 0);
INSERT INTO posts VALUES(20, 'Twentieth Post', '2024-11-05 19:30:00', 'Exploring new horizons in the twentieth post.', 'user20', 'Lifestyle', 15, 1, 0, 1);
-- Continue adding additional posts here

CREATE TABLE comments (
    id INTEGER PRIMARY KEY,
    postId INTEGER,
    createdAt DATETIME,
    author VARCHAR(20) NOT NULL,
    content TEXT,
    likesCount INTEGER,
    dislikesCount INTEGER,
    liked BOOLEAN,
    disliked BOOLEAN,
    FOREIGN KEY (postId) REFERENCES posts(id),
    FOREIGN KEY (author) REFERENCES users(username)
);

INSERT INTO comments VALUES(1, 1, '2024-11-01 08:00:00', 'user1', 'Great post! Really enjoyed reading this.', 5, 0, 1, 0);
INSERT INTO comments VALUES(2, 1, '2024-11-01 08:10:00', 'user2', 'Interesting perspective!', 2, 1, 0, 1);
INSERT INTO comments VALUES(3, 1, '2024-11-01 08:20:00', 'user3', 'I learned something new, thanks!', 3, 0, 1, 0);
INSERT INTO comments VALUES(4, 2, '2024-11-02 09:00:00', 'user4', 'Thanks for sharing this information.', 0, 1, 0, 1);
INSERT INTO comments VALUES(5, 2, '2024-11-02 09:10:00', 'user5', 'Could you elaborate on this point?', 1, 0, 1, 0);
INSERT INTO comments VALUES(6, 2, '2024-11-02 09:20:00', 'user6', 'I disagree with your opinion.', 2, 2, 0, 1);
INSERT INTO comments VALUES(7, 3, '2024-11-03 10:00:00', 'user7', 'This is a very insightful article.', 4, 1, 1, 0);
INSERT INTO comments VALUES(8, 3, '2024-11-03 10:10:00', 'user8', 'I totally agree with you.', 3, 0, 1, 0);
INSERT INTO comments VALUES(9, 3, '2024-11-03 10:20:00', 'user9', 'Could you provide more sources?', 1, 0, 0, 1);
INSERT INTO comments VALUES(10, 4, '2024-11-04 11:00:00', 'user10', 'This resonates with my experience.', 5, 0, 1, 0);
INSERT INTO comments VALUES(11, 4, '2024-11-04 11:10:00', 'user11', 'Fantastic read!', 2, 0, 1, 0);
INSERT INTO comments VALUES(12, 4, '2024-11-04 11:20:00', 'user12', 'I have a different view on this.', 1, 1, 0, 1);
INSERT INTO comments VALUES(13, 4, '2024-11-04 11:30:00', 'user13', 'Interesting, but needs more examples.', 3, 1, 0, 1);
INSERT INTO comments VALUES(14, 5, '2024-11-05 12:00:00', 'user14', 'This content is really helpful!', 3, 0, 1, 0);
INSERT INTO comments VALUES(15, 5, '2024-11-05 12:10:00', 'user15', 'Nice work!', 0, 0, 0, 0);
INSERT INTO comments VALUES(16, 5, '2024-11-05 12:20:00', 'user16', 'I appreciate your insights.', 2, 1, 0, 1);
INSERT INTO comments VALUES(17, 6, '2024-11-06 13:00:00', 'user17', 'Travel tips are always welcome!', 5, 0, 1, 0);
INSERT INTO comments VALUES(18, 6, '2024-11-06 13:10:00', 'user18', 'What a beautiful destination!', 4, 0, 1, 0);
INSERT INTO comments VALUES(19, 6, '2024-11-06 13:20:00', 'user19', 'This is just what I needed.', 1, 1, 0, 1);
INSERT INTO comments VALUES(20, 7, '2024-11-07 14:00:00', 'user1', 'A thought-provoking post!', 2, 0, 1, 0);
INSERT INTO comments VALUES(21, 7, '2024-11-07 14:10:00', 'user1', 'Your analysis is spot on.', 1, 0, 1, 0);
INSERT INTO comments VALUES(22, 7, '2024-11-07 14:20:00', 'user2', 'I found this quite boring.', 0, 2, 0, 1);
INSERT INTO comments VALUES(23, 8, '2024-11-08 15:00:00', 'user3', 'Current trends are fascinating!', 3, 0, 1, 0);
INSERT INTO comments VALUES(24, 8, '2024-11-08 15:10:00', 'user4', 'Thanks for sharing your views.', 2, 1, 0, 1);
INSERT INTO comments VALUES(25, 8, '2024-11-08 15:20:00', 'user5', 'Looking forward to more posts like this!', 1, 0, 1, 0);
INSERT INTO comments VALUES(26, 8, '2024-11-08 15:30:00', 'user5', 'You might want to check the facts again.', 0, 2, 0, 1);
-- Additional data would follow the same pattern...


CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(50) UNIQUE NOT NULL
);
INSERT INTO categories VALUES(1,'General');
INSERT INTO categories VALUES(2,'News');
INSERT INTO categories VALUES(3,'Tech');
INSERT INTO categories VALUES(4,'Lifestyle');
INSERT INTO categories VALUES(5,'Education');
INSERT INTO categories VALUES(6,'Travel');
INSERT INTO categories VALUES(7,'Health');
INSERT INTO categories VALUES(8,'Entertainment');
INSERT INTO categories VALUES(10,'Programming');
CREATE TABLE engagement (
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
INSERT INTO engagement VALUES(1,1,1,NULL,1,0);
INSERT INTO engagement VALUES(2,1,2,1,1,0);
INSERT INTO engagement VALUES(3,1,3,2,0,1);
INSERT INTO engagement VALUES(4,1,4,NULL,1,0);
INSERT INTO engagement VALUES(5,2,1,NULL,1,0);
INSERT INTO engagement VALUES(6,2,2,1,1,0);
INSERT INTO engagement VALUES(7,2,3,2,0,1);
INSERT INTO engagement VALUES(8,2,4,3,1,0);
INSERT INTO engagement VALUES(9,3,5,NULL,0,1);
INSERT INTO engagement VALUES(10,3,6,1,1,0);
INSERT INTO engagement VALUES(11,3,7,2,1,0);
INSERT INTO engagement VALUES(12,4,8,NULL,1,0);
INSERT INTO engagement VALUES(13,4,9,1,0,1);
INSERT INTO engagement VALUES(14,4,10,2,1,0);
INSERT INTO engagement VALUES(15,4,11,3,1,0);
INSERT INTO engagement VALUES(16,5,12,NULL,0,1);
INSERT INTO engagement VALUES(17,5,13,1,1,0);
INSERT INTO engagement VALUES(18,5,14,2,1,0);
INSERT INTO engagement VALUES(19,6,15,NULL,1,0);
INSERT INTO engagement VALUES(20,6,16,1,1,0);
INSERT INTO engagement VALUES(21,6,17,2,0,1);
INSERT INTO engagement VALUES(22,7,18,NULL,1,0);
INSERT INTO engagement VALUES(23,7,19,1,0,1);
INSERT INTO engagement VALUES(24,7,20,2,1,0);
INSERT INTO engagement VALUES(25,8,1,NULL,1,0);
INSERT INTO engagement VALUES(26,8,2,1,1,0);
INSERT INTO engagement VALUES(27,8,3,2,0,1);
INSERT INTO engagement VALUES(28,8,4,3,1,0);
INSERT INTO engagement VALUES(29,9,5,NULL,1,0);
INSERT INTO engagement VALUES(30,9,6,1,0,1);
INSERT INTO engagement VALUES(31,9,7,2,1,0);
INSERT INTO engagement VALUES(32,10,8,NULL,1,0);
INSERT INTO engagement VALUES(33,10,9,1,1,0);
INSERT INTO engagement VALUES(34,11,10,NULL,1,0);
INSERT INTO engagement VALUES(35,11,11,1,0,1);
INSERT INTO engagement VALUES(36,11,12,2,1,0);
INSERT INTO engagement VALUES(37,12,13,NULL,0,1);
INSERT INTO engagement VALUES(38,12,14,1,1,0);
INSERT INTO engagement VALUES(39,13,15,NULL,1,0);
INSERT INTO engagement VALUES(40,13,16,1,1,0);
INSERT INTO engagement VALUES(41,13,17,2,0,1);
INSERT INTO engagement VALUES(42,14,18,NULL,1,0);
INSERT INTO engagement VALUES(43,14,19,1,0,1);
INSERT INTO engagement VALUES(44,14,20,2,1,0);
INSERT INTO engagement VALUES(45,15,1,NULL,1,0);
INSERT INTO engagement VALUES(46,15,2,1,0,1);
INSERT INTO engagement VALUES(47,16,3,NULL,1,0);
INSERT INTO engagement VALUES(48,16,4,1,1,0);
INSERT INTO engagement VALUES(49,17,5,NULL,1,0);
INSERT INTO engagement VALUES(50,17,6,1,0,1);
INSERT INTO engagement VALUES(51,18,7,NULL,1,0);
INSERT INTO engagement VALUES(52,18,8,1,1,0);
INSERT INTO engagement VALUES(53,18,9,2,0,1);
INSERT INTO engagement VALUES(54,19,10,NULL,1,0);
INSERT INTO engagement VALUES(55,19,11,1,1,0);
INSERT INTO engagement VALUES(56,20,12,NULL,0,1);
INSERT INTO engagement VALUES(57,20,13,1,1,0);
INSERT INTO engagement VALUES(58,21,14,NULL,1,0);
INSERT INTO engagement VALUES(59,21,15,1,1,0);
INSERT INTO engagement VALUES(60,22,16,NULL,1,0);
INSERT INTO engagement VALUES(61,22,17,1,0,1);
INSERT INTO engagement VALUES(62,23,18,NULL,1,0);
INSERT INTO engagement VALUES(63,23,19,1,0,1);
INSERT INTO engagement VALUES(64,24,20,NULL,1,0);
INSERT INTO engagement VALUES(65,24,1,1,1,0);
INSERT INTO engagement VALUES(66,25,2,NULL,1,0);
INSERT INTO engagement VALUES(67,25,3,1,0,1);
INSERT INTO engagement VALUES(68,26,4,NULL,1,0);
INSERT INTO engagement VALUES(69,26,5,1,1,0);
INSERT INTO engagement VALUES(70,27,6,NULL,1,0);
INSERT INTO engagement VALUES(71,27,7,1,0,1);
INSERT INTO engagement VALUES(72,28,8,NULL,1,0);
INSERT INTO engagement VALUES(73,28,9,1,1,0);
INSERT INTO engagement VALUES(74,29,10,NULL,1,0);
INSERT INTO engagement VALUES(75,29,11,1,1,0);
INSERT INTO engagement VALUES(76,30,12,NULL,1,0);
INSERT INTO engagement VALUES(77,30,13,1,0,1);
INSERT INTO engagement VALUES(78,31,14,NULL,1,0);
INSERT INTO engagement VALUES(79,31,15,1,1,0);
INSERT INTO engagement VALUES(80,32,16,NULL,1,0);
INSERT INTO engagement VALUES(81,32,17,1,1,0);
INSERT INTO engagement VALUES(82,33,18,NULL,1,0);
INSERT INTO engagement VALUES(83,33,19,1,0,1);
INSERT INTO engagement VALUES(84,34,20,NULL,1,0);
INSERT INTO engagement VALUES(85,34,1,1,1,0);
INSERT INTO engagement VALUES(86,35,2,NULL,1,0);
INSERT INTO engagement VALUES(87,35,3,1,1,0);
INSERT INTO engagement VALUES(88,36,4,NULL,1,0);
INSERT INTO engagement VALUES(89,36,5,1,0,1);
INSERT INTO engagement VALUES(90,37,6,NULL,1,0);
INSERT INTO engagement VALUES(91,37,7,1,1,0);
INSERT INTO engagement VALUES(92,38,8,NULL,1,0);
INSERT INTO engagement VALUES(93,38,9,1,0,1);
INSERT INTO engagement VALUES(94,39,10,NULL,1,0);
INSERT INTO engagement VALUES(95,39,11,1,1,0);
INSERT INTO engagement VALUES(96,40,12,NULL,0,1);
INSERT INTO engagement VALUES(97,40,13,1,1,0);
INSERT INTO engagement VALUES(98,41,14,NULL,1,0);
INSERT INTO engagement VALUES(99,41,15,1,1,0);
INSERT INTO engagement VALUES(100,42,16,NULL,1,0);
INSERT INTO engagement VALUES(101,42,17,1,0,1);
INSERT INTO engagement VALUES(102,43,18,NULL,1,0);
INSERT INTO engagement VALUES(103,43,19,1,1,0);
INSERT INTO engagement VALUES(104,44,20,NULL,1,0);
INSERT INTO engagement VALUES(105,44,1,1,1,0);
INSERT INTO engagement VALUES(106,45,2,NULL,1,0);
INSERT INTO engagement VALUES(107,45,3,1,1,0);
INSERT INTO engagement VALUES(108,46,4,NULL,1,0);
INSERT INTO engagement VALUES(109,46,5,1,0,1);
INSERT INTO engagement VALUES(110,47,6,NULL,1,0);
INSERT INTO engagement VALUES(111,47,7,1,1,0);
INSERT INTO engagement VALUES(112,48,8,NULL,1,0);
INSERT INTO engagement VALUES(113,48,9,1,1,0);
INSERT INTO engagement VALUES(114,49,10,NULL,1,0);
INSERT INTO engagement VALUES(115,49,11,1,0,1);
INSERT INTO engagement VALUES(116,50,12,NULL,1,0);
INSERT INTO engagement VALUES(117,50,13,1,1,0);
DELETE FROM sqlite_sequence;
INSERT INTO sqlite_sequence VALUES('users',20);
INSERT INTO sqlite_sequence VALUES('posts',50);
COMMIT;
