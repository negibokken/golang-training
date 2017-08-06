# 11章 演習問題6

結果として、すべての場合で2.6.2 の結果が速い

* 下記結果との対応
    * 2.6.2 は Benchmark10 - 100000
    * 練習問題2.3は BenchmarkMyPopCount10 - 100000
    * 練習問題2.4は BenchmarkDirtyPopCount10 - 100000
    * 練習問題2.5は BenchmarkBitClearPopCount10 - 100000

```
=== RUN   TestPopCount
--- PASS: TestPopCount (0.00s)
=== RUN   TestMyPopCount
--- PASS: TestMyPopCount (0.00s)
=== RUN   TestDirtyPopCount
--- PASS: TestDirtyPopCount (0.00s)
=== RUN   TestBitClearPopCount
--- PASS: TestBitClearPopCount (0.00s)
Benchmark10-4                        	500000000	         3.61 ns/op
Benchmark100-4                       	30000000	        51.4 ns/op
Benchmark1000-4                      	 5000000	       369 ns/op
Benchmark10000-4                     	  500000	      3508 ns/op
Benchmark100000-4                    	   50000	     49773 ns/op
Benchmark1000000-4                   	    5000	    405553 ns/op
BenchmarkMyPopCount10-4              	10000000	       273 ns/op
BenchmarkMyPopCount100-4             	 1000000	      2695 ns/op
BenchmarkMyPopCount1000-4            	  100000	     21325 ns/op
BenchmarkMyPopCount10000-4           	   10000	    200104 ns/op
BenchmarkMyPopCount100000-4          	    1000	   2327368 ns/op
BenchmarkMyPopCount1000000-4         	      50	  22268669 ns/op
BenchmarkDirtyPopCount10-4           	30000000	        59.3 ns/op
BenchmarkDirtyPopCount100-4          	 2000000	       980 ns/op
BenchmarkDirtyPopCount1000-4         	  100000	     14304 ns/op
BenchmarkDirtyPopCount10000-4        	   10000	    206850 ns/op
BenchmarkDirtyPopCount100000-4       	    1000	   2319302 ns/op
BenchmarkDirtyPopCount1000000-4      	      50	  33292178 ns/op
BenchmarkBitClearPopCount10-4        	100000000	        23.6 ns/op
BenchmarkBitClearPopCount100-4       	 5000000	       343 ns/op
BenchmarkBitClearPopCount1000-4      	  300000	      4667 ns/op
BenchmarkBitClearPopCount10000-4     	   30000	     57545 ns/op
BenchmarkBitClearPopCount100000-4    	    2000	    699138 ns/op
BenchmarkBitClearPopCount1000000-4   	     200	   9292670 ns/op
PASS
ok  	github.com/negibokken/golang-training/ch11/ex06/popcount	51.577s
```
