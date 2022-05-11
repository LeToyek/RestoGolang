CREATE TABLE ORDER_FOOD(
    ORDER_ID varchar,
    FOOD_ID varchar,
    CONSTRAINT fk_order FOREIGN KEY(ORDER_ID) REFERENCES ORDER_User(Order_Id),
    CONSTRAINT pk_order primary key(ORDER_ID,FOOD_ID)
);