# gog

`gog` is a tool to execute only specific commands at executing `go generate` .

## Usage

```
$ gog [-cmd command_name] [-n] [file.go...]
Usage of gog:
  -cmd string
        command name that you want to execute
  -n    dry run
  file.go... string...
        target files of `go generate`
```

## Example

```
$ gog -cmd="echo" docs/sample1.go
[execute command]echo sample1
[results] sample1

# multi files
$ gog -cmd="echo" docs/sample1.go docs/sample2.go
[execute command]echo sample1
[results] sample1
[execute command]echo sample2
[results] sample2

# all files in a package
$ gog -cmd="echo" ./docs/...
[execute command]echo sample1
[results] sample1
[execute command]echo sample2
[results] sample2

# dry run
$ gog -cmd="echo" -n docs/sample1.go
[execute command]echo sample1
[results] sample1
```
