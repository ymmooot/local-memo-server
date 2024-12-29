Run build script:

```sh
go build -o ./dist/memo-server ./main.go
```

Create `memo.plist` file in `~/Library/LaunchAgents/` directory.

```sh 
launchctl load ~/Library/LaunchAgents/memo.plist
```

Now, the server will start automatically when you log in to your mac.  
You should bookmark the server URL in your browser.
