# zz

a little tool to get the intersection and difference between multiple files.

## Install

```shell
go get github.com/kaleocheng/zz
```

## Usage

```bash
$ cat a.txt
A
B
C
D
E
$ cat b.txt
C
D
E
F
G
```

```bash
# Get Difference between a.txt and b.txt
$ zz a.txt b.txt
A
B
F
G

# Get diff lines which only in a.txt
$ zz -only a.txt b.txt
A
B

# Get Intersection between a.txt and b.txt
$ zz -i a.txt b.txt
C
D
E
```
