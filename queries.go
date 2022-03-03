package main

const (
	createUserSql = `insert into users(
			name, login, password
		) values (
			?, ?, ?
		)`

	createTableUsersSql = `
		create table if not exists users(
			id integer not null primary key autoincrement,
			name text not null,
			login text not null unique,
			password text not null,
			timestamp datetime default current_timestamp
		)
	`

	getPasswordUserSql = "select id, login, name, password from users where login = ?"

	getUserSql = "select id, name, login, strftime('%d/%m/%Y %H:%M', timestamp) as timestamp from users where id = ?"

	getAllUsersSql = "select distinct id, name, login, strftime('%d/%m/%Y %H:%M', timestamp) as timestamp from users"
)
