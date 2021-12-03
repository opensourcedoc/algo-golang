# Data Structures and Algorithms in Go

The repo implements common data structures and algorithms in Go.

## Intro

Since Go lacks generics, more boilerplate code are needed when implementing data structures. There are several ways to implementing generics-like behaviors

- Code generation, e.g. [genny](https://github.com/cheekybits/genny)
- Use interface as in sort.Sort
- Dynamic typing, using empty interface (as C void*)

We adapt the third approach, delegating some tasks to users.

## Copyright

2017, Michelle Chen

## License

MIT
