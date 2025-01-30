This is a simple note-taking page built with Go.  
It allows you to save and automatically restore your notes, even after closing and reopening the browser tab.

## Features

This app saves your notes to the browser's local storage.  
So, you can not access your notes from another browser or device.

## Setup

Build:

```sh
go build -o ./dist/memo-server ./main.go
```

Create `memo.plist` file in `~/Library/LaunchAgents/` directory.

```sh 
launchctl load ~/Library/LaunchAgents/memo.plist
```

Now, the server will start automatically when you log in to your mac.  
You should bookmark the server URL in your browser.
