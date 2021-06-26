## Basic comparison between Go, Vala, C and C++

Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz

## Recursive fibonacci

```
$ gcc -Wall fib.c -o cfib
$ time ./cfib 45
1134903170
real    0m6,033s
user    0m6,008s
sys     0m0,017s
```

```
$ valac fib.vala -o vfib
$ time ./vfib 45
1134903170

real    0m7,079s
user    0m7,032s
sys     0m0,017s
```

```
$ go build gfib.go
$ time ./gfib 45
1134903170

real    0m5,867s
user    0m5,861s
sys     0m0,012s
```

```
$ g++ fib.cpp -o cppfib
$ time ./cppfib 45
1134903170
real    0m5,993s
user    0m5,990s
sys     0m0,000s
```