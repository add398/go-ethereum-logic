

ART  

|  number of address |  time    | space  |
| ----  |  ----           | ---- |
| 1e    | 108838243500 ns/op | 19893734656 B/op  |
| 8kw   | 82159467400 ns/op	 | 15999380656 B/op  |
| 6kw   | 60167254800 ns/op	 | 11934526144 B/op  |
| 4kw   | 35961011000 ns/op	 | 7849864160 B/op   |
| 2kw   | 15677027900 ns/op	 | 3893014432 B/op   |
| 1kw   | 7518902400 ns/op	 | 1981318856 B/op   |
| 1e value 1byte   | 88371681200 ns/op	 | 17493732960 B/op   |


MPT

|  number of address |  time    | space  |
| ----  |  ----           | ---- |
| 1e    | 108838243500 ns/op | 19893734656 B/op  |



20byte key   100byte value   

100w        53.1MB  543MB   1.9G

Benchmark_3_10000w-12    	       1	1678151900 ns/op	1930750848 B/op	10653398 allocs/op

1000w       486MB   5342MB  21G

Benchmark_3_10000w-12    	       1	21971464100 ns/op	21972485784 B/op	114685602 allocs/op

2000w       967     11250MB  45GB        

Benchmark_3_10000w-12    	       1	49683473800 ns/op	45557796512 B/op	235070902 allocs/op

3000w       1448  17413MB   69

Benchmark_3_10000w-12    	       1	82421237300 ns/op	69864753808 B/op	357834385 allocs/op


4000w    2409 (5kw 的 slice)  21393MB   94G

Benchmark_3_10000w-12    	       1	111514756800 ns/op	94588800400 B/op	481625021 allocs/op



交易数量  一次 commit 5000

100w   78.6MB   508MB

500w   377MB   2283MB

1000w   754 MB   4592MB

2000w   1504MB   9093MB

3000w   1985MB   15031MB

4000w   2643MB   18991MB

5000w   3302MB   22700MB


单次 commit 的数量影响最后的大小， 例子里面大多是几千个一起commit



6000w   78.6MB   508MB





