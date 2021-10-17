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

create table cathegory (
    id serial,
	cathegory text
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
	is_active boolean DEFAULT true
)

create table users_id_cathegory_id (
    id serial,
	users_id integer,
	cathegory_id integer
)

create table keys (
    id serial,
	users_id integer,
	api_key text
)