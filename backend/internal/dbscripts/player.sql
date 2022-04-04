create table if not exists artists (
	id bigserial primary key,
	name varchar,
    storageFolderName varchar,
)

create table if not exists albums (
	id bigserial primary key,
	name varchar,
    genre varchar,
    trackIDs integer[],
    artistID bigserial references artists(id)
)

create table if not exists tracks (
	id bigserial primary key,
	name varchar,
    fullName varchar,
    storageID varchar,
	artistID bigserial references artists(id),
    albumID bigserial references albums(id)
)