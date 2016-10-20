# wp

Increase one's word power.

## Example

```
$ touch some-sentenece
$ cat <<< EOS
Go (often referred to as golang) is a free and open source programming language created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing, memory safety features and CSP-style concurrent programming features added.

The language was announced in November 2009; it is used in some of Google's production systems, as well as by other firms. Two major implementations exist: Google's Go compiler, "gc", is developed as open source software and targets various platforms including Linux, OS X, Windows, various BSD and Unix versions, and since 2015 also mobile devices, including smartphones. A second compiler, gccgo, is a GCC frontend. The "gc" toolchain is self-hosting since version 1.5.

EOS
$ wp search -i -f some-sentence -w open
```
