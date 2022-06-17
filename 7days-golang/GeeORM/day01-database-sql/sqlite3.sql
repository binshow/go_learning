sqlite3 gee.db  # 创建数据库

CREATE TABLE User(Name text, Age integer);
INSERT INTO User(Name, Age) VALUES ("Tom", 18), ("Jack", 25);
.head on
SELECT * FROM User WHERE Age > 20;
SELECT COUNT(*) FROM User;
.table