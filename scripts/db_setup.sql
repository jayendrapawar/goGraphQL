CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE follows (
    follower_id INT,
    following_id INT,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (follower_id, following_id)
);

INSERT INTO users (name) VALUES ('1');
INSERT INTO users (name) VALUES ('2');
INSERT INTO users (name) VALUES ('3');
INSERT INTO users (name) VALUES ('4');
INSERT INTO users (name) VALUES ('5');

