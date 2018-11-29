# readpe


## What is readpe

readpe is a parsing tool for PE binary such as ```readelf```.

## How to use

```
$ go get github.com/owlinux1000/readpe
$ readpe -h
PE Parser like readelf

Usage: 
	readpe [option] [<args>]

Options:
	-d		Display the MS-DOS header
	-f		Display file header
	-h		Display this message
	-o		Display the optional header
	-s		Display section headers
	-v		Print version info and exit
```

## Feature

- Support for some structure of PE Header
  - IMAGE\_DOS\_HEADER
  - IMAGE\_NT\_HEADER
  - IMAGE\_FILE\_HEADER
  - IMAGE\_OPTIONAL_HEADER
  - IMAGE\_SECTION\_HEADER
