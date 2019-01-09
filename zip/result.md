# Сохранение разультатов парсинга

## !!! в самом приложении, map, без разделения для likes

Alloc = 1544 MiB	TotalAlloc = 2694 MiB	Sys = 1693 MiB	Free = 49 MiB	NumGC = 15
Readed file with number: 130
Alloc = 1201 MiB	TotalAlloc = 2694 MiB	Sys = 1693 MiB	NumGC = 16
Ready for requests...

### Тоже самое, но без индексов email и phone

Alloc = 1382 MiB	TotalAlloc = 2498 MiB	Sys = 1492 MiB	Free = 50 MiB	NumGC = 15
Readed file with number: 130
Alloc = 1086 MiB	TotalAlloc = 2498 MiB	Sys = 1493 MiB	NumGC = 16

## В единый map[int]Account с likes внутри

Alloc = 1302 MiB	TotalAlloc = 2181 MiB	Sys = 1425 MiB	NumGC = 14
Readed file with number: 130
After:
Alloc = 1085 MiB	TotalAlloc = 2181 MiB	Sys = 1425 MiB	NumGC = 15

## В 2 map для аккаунтов и для лайков 

Alloc = 1336 MiB	TotalAlloc = 2278 MiB	Sys = 1490 MiB	NumGC = 14
Readed file with number: 130
After:
Alloc = 1110 MiB	TotalAlloc = 2278 MiB	Sys = 1490 MiB	NumGC = 15

## В заранее не предопределенны slice

Alloc = 1506 MiB	TotalAlloc = 3414 MiB	Sys = 2010 MiB	NumGC = 17
Readed file with number: 130
After:
Alloc = 1061 MiB	TotalAlloc = 3414 MiB	Sys = 2010 MiB	NumGC = 18

## В предопредленный slice размером 1300000

Alloc = 1194 MiB	TotalAlloc = 2107 MiB	Sys = 1486 MiB	NumGC = 6
Readed file with number: 130
After:
Alloc = 1046 MiB	TotalAlloc = 2107 MiB	Sys = 1487 MiB	NumGC = 7

## Чистая запись (не очень много данных), без чтения json и геренации мусора

Alloc = 824 MiB	TotalAlloc = 972 MiB	Sys = 880 MiB	NumGC = 4
Readed file with number: 130
After:
Alloc = 773 MiB	TotalAlloc = 972 MiB	Sys = 881 MiB	NumGC = 5

### Тоже самое, но в map

Alloc = 889 MiB	TotalAlloc = 1035 MiB	Sys = 953 MiB	NumGC = 10
Readed file with number: 130
After:
Alloc = 803 MiB	TotalAlloc = 1035 MiB	Sys = 954 MiB	NumGC = 11

## Чистая запись (болле или менее много данных), без чтения json и геренации мусора

Alloc = 1180 MiB	TotalAlloc = 1309 MiB	Sys = 1284 MiB	NumGC = 4
Readed file with number: 130
After:
Alloc = 1051 MiB	TotalAlloc = 1309 MiB	Sys = 1286 MiB	NumGC = 5

### Тоже самое, но в map

Alloc = 1194 MiB	TotalAlloc = 1373 MiB	Sys = 1291 MiB	NumGC = 10
Readed file with number: 130
After:
Alloc = 1081 MiB	TotalAlloc = 1373 MiB	Sys = 1293 MiB	NumGC = 11

## С чтением архива (прилично данных), итореированием по нему, но без парсинга и записаю в предопределенны slice на 1300000 с лайками внутри

Alloc = 1220 MiB	TotalAlloc = 1399 MiB	Sys = 1351 MiB	NumGC = 4
Readed file with number: 130
After:
Alloc = 1083 MiB	TotalAlloc = 1399 MiB	Sys = 1352 MiB	NumGC = 5
