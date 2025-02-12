-- Create users table
CREATE TABLE users (
                       userid VARCHAR(255) PRIMARY KEY,
                       username VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'Asia/Almaty')
);

-- Create feedbacks table
CREATE TABLE feedbacks (
                           feedback_id SERIAL PRIMARY KEY,
                           userid VARCHAR(255) NOT NULL,
                           feedback TEXT NOT NULL,
                           created_at TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'Asia/Almaty'),
                           CONSTRAINT fk_user
                               FOREIGN KEY(userid)
                                   REFERENCES users(userid)
                                   ON DELETE CASCADE
);
