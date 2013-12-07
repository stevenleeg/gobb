# gobb
A simple forum platform written in Go. 

## Work in progress
This project is waaay not ready for the public yet. I only put it up so early because one of my friends was interested in seeing what was causing me to pester him with Go questions so much (thanks [Aditya](http://github.com/chimeracoder)!).

Feel free to browse around and check back, but this project is not ready yet!


### Installation

````sh
$ go get github.com/stevenleeg/gobb/gobb
````

Copy gobb.sample.conf into gobb.conf and set the `[database]` parameters.

Next, set up Postgres as follows:
```
$ psql
# CREATE DATABASE gobb;
CREATE DATABASE
# CREATE ROLE gobb WITH PASSWORD 'password';
CREATE ROLE
# GRANT ALL PRIVILEGES ON DATABASE gobb TO gobb;
GRANT
# ALTER ROLE gobb WITH LOGIN;
ALTER ROLE
# ^D
~$ psql -h localhost -U gobb -W gobb
gobb=> \i your/gopath/src/github.com/stevenleeg/gobb/models/schema.sql 
```

Add a board to start:
```
$ psql gobb
# INSERT INTO boards (title, description) VALUES('general', 'whatever');
```

Run!
```
$ go run gobb/main.go
```