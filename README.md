# sgogen

`sgogen` is a tool to execute only specific commands at executing `go generate` .

## Usage

```
$ sgogen [-cmd command_name] [-n] [file.go...]
Usage of sgogen:
  -cmd string
        command name that you want to execute
  -n    dry run
  file.go... string...
        target files of `go generate`
```

## Example

```
$ sgogen -cmd="echo" docs/sample1.go
[execute command]echo sample1
[results] sample1

# multi files
$ sgogen -cmd="echo" docs/sample1.go docs/sample2.go
[execute command]echo sample1
[results] sample1
[execute command]echo sample2
[results] sample2

# all files in a package
$ sgogen -cmd="echo" ./docs/...
[execute command]echo sample1
[results] sample1
[execute command]echo sample2
[results] sample2

# dry run
$ sgogen -cmd="echo" -n docs/sample1.go
[execute command]echo sample1
[results] sample1
```
