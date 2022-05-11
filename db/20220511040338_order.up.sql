CREATE TABLE ORDER_User(
    Order_Id   varchar primary key,   
	Order_date varchar,   
	Created_at varchar,   
	Updated_at varchar,   
	User_id    varchar,   
	Table_id   int,
    CONSTRAINT fk_user FOREIGN KEY(User_id) REFERENCES users(User_id),
    CONSTRAINT fk_table FOREIGN KEY(Table_id) REFERENCES table_resto(Number)   
);
