allfmt
======

One entry point to all of your fmt tools like jsfmt, gofmt, dartfmt, etc'. This will call the matching `fmt` from your IDE of choice provided that **you have already installed jsfmt, gofmt, dartfmt etc'**


How to enable code formatting for dart, golang and more in Intellij/WebStorm
============================================================================


`$ go get github.com/amitit/allfmt`

Check if this was correctly installed:

```
$ allfmt
Error: --input must be specified
Usage: allfmt [global options]

Global options:
        -d, --dart  Path to dartfmt (default: dartfmt)
        -g, --go    Path to gofmt (default: gofmt)
        -j, --js    Path to jsfmt (default: jsfmt)
        -i, --input Input source (*)
```

Typically this will get executed from your `GOPATH/bin`

Now in IntelliJ: Preferences -> IDE Settings -> External Tools create a new tool:

Fill in the dialog:
```
Name: allfmt
Group: Formatting tools
Check-boxes: Synchronize files after execution
Program: ./allfmt
Parameters: --input $FilePath$
Working directory: <your GOPATH>/bin 
```

Now let's configure a key combination for this:
----------------------------------------------

In IntelliJK go to Preferences -> IDE Settings -> Keymap -> External Tools 
-> Formatting Tools -> allfmt (double click)
-> add keyboard shortcut

I chose alt-cmd-K because it's close to alt-cmd-L which is default Intellij for code formatting, but you can choose everything you want.

