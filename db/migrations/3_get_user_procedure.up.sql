CREATE PROCEDURE spGetUsers()
BEGIN
   SELECT id, name, first_name, last_name, email, phone, age FROM `user`;
END