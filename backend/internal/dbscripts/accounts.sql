create table if not exists accounts (
	id bigserial primary key,
	firstName varchar,
	lastName varchar,
	username varchar,
	email varchar not null unique,
	password varchar not null
)

create table if not exists playlists (
	id bigserial primary key,
	name varchar,
	trackIDs integer[],
	accountID BIGSERIAL references accounts(id)
)