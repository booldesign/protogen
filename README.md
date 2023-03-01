# protogen

```bash gen.sh```
生成更新后的kitex_gen代码后提交打tag即可
后期交由idl仓库合码后自动更新


# 生成protogen
```
# bd @ bddeMacBook-Pro-4 in ~/go/src/tmp/account [15:42:13] C:1
$ cd ../protogen

# bd @ bddeMacBook-Pro-4 in ~/go/src/tmp/protogen [15:44:58]
$ go mod init example.com/protogen
go: creating new go.mod: module example.com/protogen

# bd @ bddeMacBook-Pro-4 in ~/go/src/tmp/protogen [15:45:07]
$ bash gen.sh                            
ls: ../protos/*/*/*/service.proto: No such file or directory

# bd @ bddeMacBook-Pro-4 in ~/go/src/tmp/protogen [15:45:39]
$ ls ../protos/*/*/*/service.proto
zsh: no matches found: ../protos/*/*/*/service.proto

# bd @ bddeMacBook-Pro-4 in ~/go/src/tmp/protogen [15:47:06] C:1
$ ls ../protos/*/*/service.proto
../protos/app/account/service.proto

# bd @ bddeMacBook-Pro-4 in ~/go/src/tmp/protogen [15:47:31]
$ bash gen.sh 
```