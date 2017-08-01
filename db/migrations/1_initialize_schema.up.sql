create table posts (
	id 			bigserial primary key,
	content 	text,
	created_at 	timestamp default now()
);
