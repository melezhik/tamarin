#!/usr/bin/env tamarin

conn := sql.connect("postgres://postgres:mysecretpassword@localhost:5432/postgres")
result := sql.query(conn, "select * from users where age > $1", 40)
switch result.is_ok() {
    case true:
        rows := result.unwrap()
        rows.each(func(row) { print("row:", row) })
        print(len(rows))
        print(type(rows))
    case false:
        print("Failed:", result)
}

conn | sql.query("select COUNT(*), NOW() from users")