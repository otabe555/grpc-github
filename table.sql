create table contactlist (
    id serial PRIMARY KEY,
    name varchar(80) not null,
	phone varchar(13) not null,
	gender varchar(6) not null,
	email varchar(50) not null,
	createat timestamp default now()
);
create table tasklist(
	id serial primary key,
	assignee text not null,
	title text not null,
	deadline timestamp not null,
	done boolean
);