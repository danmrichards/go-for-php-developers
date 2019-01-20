# Go for PHP Developers
In this repo are two applications providing identical REST APIs. One implementation
is in PHP and the other in Golang.

The purpose of this is to benchmark the two applications with the goal of illustrating
the performance difference.

Both applications are in realistic web stacks, provided by docker, with identically
configured nginx reverse proxies. The Golang version can easily run without it
but the PHP built in web-server is so slow that it would make the benchmark pointless.

## Usage
See the README.md files in the go and php folders.

## Benchmark Results
Below are the benchmark results I see on my machine:

```
TBC
```