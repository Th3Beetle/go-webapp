create table earnings (
    id serial,
	users_id integer,
	amount numeric,
	currency_id integer,
	cathegory_id integer,
	ear_date timestamp,
	ear_comment text
)

create table expenses (
    id serial,
	users_id integer,
	amount numeric,
	currency_id integer,
	cathegory_id integer,
	exp_date timestamp,
	exp_comment text
)

create table earnings (
    id serial,
	cathegory text
)

create table earnings (
    id serial,
	name text
)

create table user_data (
    id serial,
	name text,
	surname text,
	phone text,
	email text
)

create table users (
    id serial,
	login text,
	password text,
	date_created timestamp,
	is_active boolean
)

create table users_id_cathegory_id (
    id serial,
	users_id integer,
	cathegory_id integer
)