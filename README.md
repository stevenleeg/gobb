# GoBB
A simple forum platform written in Go. 

## Warning! Alpha quality software!
GoBB is currently in its early stages of development. While it is pretty usable at the moment (and actively being used in a *trusted* production environment), I'd reccommend you hold off on using it for something big for the time being. There are still a lot of things that need to be patched up before it's ready for the big leagues. Having said that, if you're looking for a simple (and blazing fast) bulletin board for your friends who are willing to deal with some bugs, this might be the bb for you!

GoBB is getting better by the day, so hopefully it'll be ready to graduate from the alpha stage of development soon.

## Installation

1. Get GoBB
````sh
$ go get github.com/stevenleeg/gobb/gobb
````

2. Copy gobb.sample.conf into gobb.conf and set the `[database]` parameters.

3. Set up the Postgres database:

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
```

3. Run!
The first time you run gobb you'll need to pass the `--migrate` flag. This will automatically set up the database schema for you.

```
$ gobb --config /path/to/gobb.conf --migrate
```

The server will be up and running on port 8080, where you can create your first account.

4. Admin yourself
After creating your account, admin yourself so you can create boards and moderate posts.

```
$ psql gobb
# UPDATE users SET group_id='1' WHERE id='1';
UPDATE 0 1
```

And that's it! You should have a functional copy of GoBB ready to use!

If you understand what you're getting yourself into and willing to run GoBB in a prod environment, I reccommend setting up an nginx reverse-proxy to expose your installation to the public. Create a new nginx config that looks something like this:

```Nginx
server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```
