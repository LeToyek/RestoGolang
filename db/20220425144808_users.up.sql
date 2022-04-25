CREATE TABLE users(
    ID            serial primary key,
	First_name    varchar,
	Last_name     varchar,
	Email         varchar,
	Password      varchar,
	Phone         varchar,
	Token         varchar,
	Refresh_Token varchar,
	Created_at    varchar,
	Updated_at    varchar,
	User_id       varchar
);