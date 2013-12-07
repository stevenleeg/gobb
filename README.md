# gobb
A simple forum platform written in Go. 

## Work in progress
This project is waaay not ready for the public yet. I only put it up so early because one of my friends was interested in seeing what was causing me to pester him with Go questions so much (thanks [Aditya](http://github.com/chimeracoder)!).

Feel free to browse around and check back, but this project is not ready yet!


### Installation

````sh
$ go get github.com/stevenleeg/gobb/gobb
$ go get bitbucket.org/liamstask/goose/cmd/goose #for migrations
````

1. Copy gobb.sample.conf into gobb.conf and set the `[database]` parameters.

2. Set up the Postgres database:
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
```

3. Run migrations using Goose:
	```
	$ goose up
	```

4. Add a board to start:
	```
$ psql gobb
# INSERT INTO boards (title, description) VALUES('general', 'whatever');
INSERT 0 1
# ^D
```

5. Run!
	```
$ cd gobb
$ go run main.go
```