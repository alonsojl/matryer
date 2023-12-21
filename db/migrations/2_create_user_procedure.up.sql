CREATE PROCEDURE spCreateUser(
 IN _name       VARCHAR(20),
 IN _first_name VARCHAR(30),
 IN _last_name  VARCHAR(30),
 IN _email      VARCHAR(30),
 IN _phone      VARCHAR(10),
 IN _age        TINYINT(1)
)
BEGIN
   INSERT INTO `user`(name, first_name, last_name, email, phone, age)VALUES(_name, _first_name, _last_name, _email, _phone, _age);
   SELECT LAST_INSERT_ID();
END