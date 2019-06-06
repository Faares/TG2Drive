# TG2Drive

Telegram Bot for uploading files to Google Drive

# Usage

first setup configuration file, in JSON foramt:
```
{
    "telegram":
    {
        "token":"BOT TOKEN", 
        "authorized":{ // users can contact the bot, unlisted users will be ignored
            "UserName#1":{ // "UserName#1" telegram username
                "name":"Name#1", // using in replay message
                "password":"Password", // to login
                "user":"User#1"    // in case we need it in future
            },
            "UserName#2":{ // another user
                "name":"Name#2",
                "password":"Passowrd",
                "user":"User#2"
            }
        }
    },

    "drive":{
        // DRIVE CONFIG, download it from Google Developers Console
    }
      
}
```
then:
```bash
go install
TG2Drive -config=/path/to/config.json
```

# Development
create `Config.json` in work directory, and then:
```bash
./dev
```


have fun ;)