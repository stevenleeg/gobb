# GoBB
A simple forum platform written in Go. 

## Warning! Alpha quality software!
GoBB is currently in its early stages of development. While it is pretty usable at the moment (and actively being used in a *trusted* production environment), I'd recommend you hold off on using it for something big for the time being. There are still a lot of things that need to be patched up before it's ready for the big leagues. Having said that, if you're looking for a simple (and blazing fast) bulletin board for your friends who are willing to deal with some bugs, this might be the bb for you!

GoBB is getting better by the day, so hopefully it'll be ready to graduate from the alpha stage of development soon.

A demo of GoBB is up and running [here](http://gobb.stevegattuso.me/). Try it out!

![Screenshot](http://imgur.com/JFSzskx.png)

## Installation

### Get GoBB
````sh
$ go get github.com/stevenleeg/gobb/gobb
````

### Create a config file
Create a directory where you'd like to house your GoBB installation's files. This can be anywhere on your filesystem. Then copy the contents of [gobb.sample.conf](https://github.com/stevenleeg/gobb/blob/master/gobb/gobb.sample.conf) into a file called `gobb.conf`.

The config should be self explanatory, so fill it with the appropriate information and save.

### Set up the Postgres database:
We'll need to create a new database for GoBB's data to be stored:

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

### Run it
Now that everything's ready to go we can start the server for the first time. You'll need to pass the `--migrate` flag on the initial start up. This will automatically create the database schema and make sure it's up to date. 

```
$ gobb --config /path/to/gobb.conf --migrate
```

(Remember to add `$GOPATH/bin` to your `$PATH`! If you don't know what this means, [start here](http://askubuntu.com/questions/60218/how-to-add-a-directory-to-my-path))

The server should then be up and running on port 8080.

### Your first account
Once the server is up, go ahead and browse to http://localhost:8080 and register. Once you've created the first account, it will be promoted to admin so you can create the first boards and begin moderating posts.

And that's it! You should have a functional copy of GoBB ready to use!

### Deploy it!
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

## Contributing
If you see something that could be better with GoBB, feel free to fork it and create a pull request. Just be sure to run `go fmt` before you submit the pull request so everything stays tidy!
