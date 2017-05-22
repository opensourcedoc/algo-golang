# Data Structures and Algorithms in Go

The repo implements common data structures and algorithms in Go.

## Intro

Since Go lacks generics, more boilerplate code are needed when implementing data structures. There are several ways to implementing generics-like behaviors

- Code generation: use templates, e.g. [genny](https://github.com/cheekybits/genny)
- Use interface: as in Sort
- Dynamic typing: use empty interface as C void*

We adapt the third approach, delegating some tasks to users.

## Copyright

2017, Michael Chen

## License

MIT
