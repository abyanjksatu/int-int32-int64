# int vs int32 vs int64

Here we compare between int vs int32 vs int64 in golang.

```bash
$ go run main.go
Unsafe Size int32
4
Unsafe Size int64
8
Unsafe Size int
8

int32
2147483647
-2147483648
 
 
int64
9223372036854775807
-9223372036854775808
 
 
int
9223372036854775807
-9223372036854775808
 
 
MaxInt == MaxInt64
true
```

As we can see max of int == int64 is true. which mean same max value

source: https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go