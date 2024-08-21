### Упражнение 4.5

Напишите функцию, которая без выделения дополнительной памяти удаляет все смежные дубликаты в срезе []string. 

Решение:

https://github.com/unixlinuxgeek/gopl/blob/db35830e1497be2da8dba44c3e1abe263bb22a7d/strutil/strutil.go#L1-L26

Модульные тесты:
```
$ go test -v
=== RUN   TestRemDup
Input:  [a a b c d d d e m q q] Capacity: 11
Output: [a b c d e m q]         Capacity: 11
--- PASS: TestRemDup (0.00s)
PASS
ok      strutil 0.002s
```
