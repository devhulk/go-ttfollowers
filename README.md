# go-ttfollowers

Simple tiktok followers bot. 

1. Get list of creators from repo
2. Pull data from secret API
3. Output followers, likes, video count
4. Mechanism for adding users.


## Run Instructions

Create env var in terminal that maps the name of the user config file you want to load.
```
export USER_FILE_NAME="userlist.json"
```

Create env var for top secret url.
```
export TT_SECRET_URL="<INSERT YOUR TOP SECRET URL HERE>"
```

Run project with the following command.
```
go run *.go
```
