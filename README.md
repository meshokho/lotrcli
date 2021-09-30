
# Introduction
---
Hello! This is my pet project. I am learning to use Cobra lib for Golang (https://github.com/spf13/cobra).

This project is an CLI app which gives you some information about Lord Of The Rings universe using The One API (https://the-one-api.dev/)

---

# Installation
---

`$ go get github.com/meshokho/lotrcli`

If you get an error during the installation like this:
`cannot find package "github.com/hashicorp/hcl/hcl/printer"`
Then try to run "go get" with flags like:
```
// Go 1.15 and below
$ GO111MODULE=on go get -u -t github.com/meshokho/lotrcli

// Go 1.16 only
go install github.com/meshokho/lotrcli
```

---
# Usage
---

v. 0.1 of this app have only two commands:

### 1. phrase
`lotrcli phrase`
Fetches a random phrase from LOTR books or movies.

```
lotrcli phrase
I do not believe this darkness will endure.
```

### 2. character
`lotrcli character --name=characterName`
Fetches information about character from LOTR (Birth, race, realm, etc.)
If *--name* flag is not specified, info about Gandalf will be shown.
Do not carry about upper- or lower-case!
Also the character's name does not have to be written in full. The command will provide all suitable options. For example,
`lotrcli character --name=baggins`
will show you information about all 40 Baggins from LOTR.

---

# ToDo
---

- Add an option for **phrase** command to fetch a phrase of a custom character
- Add more comands to learn about other Cobra's features
- Add authorization for access The One API

---
# License
lotrcli is released under the Apache 2.0 license.