# urlcheck

CLI tool for checking the http responses for a given list of URLs.

There are curently 2 ways of using `urlcheck` the standard way and the "inline" way (using the -s option) that return an exit code > 0 in case of error, a timeout or an http status code < 500 

Usage:

```sh
$ urlcheck https://go.dev/ https://charm.sh/
$ urlcheck -f ~/urls.txt
```

## Inline and script usage

Check if go.dev is responding and wait max 20 seconds for the response:

```sh
#!/bin/bash 

if urlcheck -s http://go.dev -t 20 1>/dev/null 2>&1; then
  echo go.dev is UP
else
  echo go.dev is DOWN
fi
```

## Dependencies

* [bubble tea](https://github.com/charmbracelet/bubbletea/tree/master/tutorials/commands/)

* [cobra](https://cobra.dev/)

## Todos 

- [ ] use timeoutRoundTripper instead of http.Cliet.Get also for urlist checking in urlcheck.go
- [ ] -s flag : check also for the whole urilist and exit at the first fail
- [ ] feat: continuous mode with retry interval


