CREATE TABLE users (
                       uid bigserial not null primary key,
                       email varchar(320) not null,
                       password varchar(100) not null,
                       nickname varchar(20)
)