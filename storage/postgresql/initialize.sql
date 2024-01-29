CREATE TABLE users (
   uid bigserial not null primary key,
   email varchar(320) not null,
   password varchar(64) not null,
   nickname varchar(20)
);

CREATE TABLE user_role (
    uid bigserial not null references users(uid) on delete cascade,
    role varchar(20) not null,
    primary key (uid, role)
);