# Image storage telegram bot
Telegram bot that can be used as image storage

## Settings
1. Create storage directory
- storage:
    - dir1:
        - img1.jpg
        - img2.png
    - dir2:
        - img1.jpg

2. Create telegram bot
3. Add `.env` file in root 
```
BOT_TOKEN=[YOUR TELEGRAM BOT TOKEN]
STORAGE_PATH=[ABSOLUTE PATH TO STORAGE DIR]
```

## Run
`go run cmd/tgbotimgstor/main.go`
