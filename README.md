# Sheddit

![Sheddit](https://i.imgur.com/ZZbe5cW.png)

A Go REST API to schedule posts for reddit.


## Setup

There are a few configurations that need to be done.

 **NOTE: this is a script type app for reddit, so it has only access to developers account, you can add more accounts from reddit app creation board**
 
1. Go to reddit's [app creation page](https://ssl.reddit.com/prefs/apps)
2. Create an app, remember to select **script** from the radio button menu.
3. Note down client secret and client id (both will appear in the app details once it is created)
4. Create .env file in the root of this projects directory (sample env file below)
5. Run the program with the command `godotenv -f .env go run main.go`
6. Now you can send post requests to **/schedulePost** with the following data
```
      "subreddits":"test", (In case multiple, comma seperate them)
      "title":"title-of-your-post",
      "link":"any-validurl",
      "scheduledate":"2020-03-06T10:46:00.000Z" (ISO TimeDate String)
```
Currently only supports link posts.

Sample .env file looks like this

```
  clientid=your-client-id-goes-here
  clientsecret=your-client-secret-goes-here
  rediretcurl=this-can-be-anything-example-http://localhost:8000
  useragent=sample-useragent
  username=your-reddit-username
  password=your-reddit-password
  auth_url=www.reddit.com
  secure_api=https://oauth.reddit.com/api
  bucketname=any-string-is-fine
```
*A bucket (analogous to a db table) will be created with bucketname*

## Roadmap

Things that will be supported soon

* Text Posts
* Status of scheduled posts

## Contact me?

DM me on [twitter](https://twitter.com/rahulnpadalkar). If you like my work and want to support me then [buy me a coffee](https://www.buymeacoffee.com/1UyiBMG)!
