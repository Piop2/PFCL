# PFCL
> Pio’s Freakin’ Cool Language

## What is it?
PFCL (Pretty Functional Config Language)  
A Go package for encoding and decoding a custom config language.  
Made for learning Go and the State Pattern.

## Install
```bash
go get github.com/Piop2/PFCL@main
```

## Example
```
# comment
title = "My Config"
version = "1.0.0"
enabled = true

list  = {"a", "b", "c"}
range = (1 .. 5)

[table]
name    = "users"
rows    = 100
active  = true

[table.user]
id   = 1
name = "Alice"
```

> I know it looks like TOML. I don’t care. it’s just for learning.