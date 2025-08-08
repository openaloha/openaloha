# quick start


# modules
## openaloha-sidecar
```shell
go run main.go -workspace="/Users/Shared/Projects/test/openaloha-test" -sync.type="git" -sync.git.url="https://github.com/openaloha/openaloha-test.git" -sync.git.branch="main" -sync.git.syncInterval="10s" -run.init.cmds="cd ./openaloha-front-test/openaloha-front-vue-test && npm install && npm run dev" -run.refresh='[{"condition":"*.vue","cmds":[]}]'
```