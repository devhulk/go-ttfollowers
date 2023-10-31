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

## Current Slack Integration Method

Currently, mostly because I wanted to mess with Go templates, I am using a template to create the slack json request.

The template is ```slack.tmpl.json```. Don't edit this file unless absolutely necessary.
The actual request that is being sent to slack is in ```slack_request.json```. 
You can copy this generated file into [Block Kit Builder](https://app.slack.com/block-kit-builder) to trouble shoot errors.
