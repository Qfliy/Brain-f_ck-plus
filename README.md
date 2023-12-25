# Brain f*ck plus

Brain f*ck plus (bf+) - realization [brain f*ck](https://en.wikipedia.org/wiki/Brainfuck#References) on [golang](https://en.wikipedia.org/wiki/Go_(programming_language)) v: 2.0

## Commands

```bin
+ - increment address in pointer
- - decrement address in pointer
> - increment pointer
< - decrement pointer
[ - begin loop
] - end loop
. - putchar
, - getchar
else - comment

* in sell & - exit
```

## Program flags

```bin
defolt - info
eg:
  > bf+
  ...

str - intrepretate program as string
eg:
  > bf+ str ">++++++++[<+++++++++>-]<."
  H

ran - intrepretate program fail
eg:
  > bf+ run "..\programs\test1.bf"
  Hello world!

shell - ran shell
eg:
  >bf+ shell
  ?> &

  >
```

## Autor

24.12.2023
[Qfli!](https://github.com/Qfliy)
Heppy new year :)
