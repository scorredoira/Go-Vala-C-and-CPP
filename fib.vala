void main(string[] args) {
    int i = int.parse(args[1]);
    print("%d\n", fib(i));
}

int fib(int i) {
    if(i < 2) {
        return i;
    }
    return fib(i - 1) + fib(i - 2);
}