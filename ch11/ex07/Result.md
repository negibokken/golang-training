# 11章 練習問題7

* 全体の傾向として、 32, 64, map実装の順に速い

```
BenchmarkHas100-4           	  500000	      3632 ns/op
BenchmarkHas1000-4          	   50000	     39626 ns/op
BenchmarkHas10000-4         	    5000	    378769 ns/op
BenchmarkAdd100-4           	  300000	      4369 ns/op
BenchmarkAdd1000-4          	   30000	     43741 ns/op
BenchmarkAdd10000-4         	    3000	    428135 ns/op
BenchmarkAddAll100-4        	   30000	     55457 ns/op
BenchmarkAddAll1000-4       	     300	   5156412 ns/op
BenchmarkAddAll10000-4      	       2	 594993779 ns/op
BenchmarkUnionWith100-4     	 1000000	      2448 ns/op
BenchmarkUnionWith1000-4    	 1000000	      2356 ns/op
BenchmarkUnionWith10000-4   	  500000	      2286 ns/op
BenchmarkString100-4        	     200	   8822679 ns/op
BenchmarkString1000-4       	     100	  10388534 ns/op
BenchmarkString10000-4      	     100	  10064531 ns/op
PASS
ok  	github.com/negibokken/golang-training/ch11/ex07/intset	27.470s
```

```
BenchmarkHas100-4           	  500000	      3645 ns/op
BenchmarkHas1000-4          	   50000	     34502 ns/op
BenchmarkHas10000-4         	    5000	    347589 ns/op
BenchmarkAdd100-4           	  500000	      3869 ns/op
BenchmarkAdd1000-4          	   50000	     37976 ns/op
BenchmarkAdd10000-4         	    5000	    384307 ns/op
BenchmarkAddAll100-4        	   30000	     52285 ns/op
BenchmarkAddAll1000-4       	     300	   4960867 ns/op
BenchmarkAddAll10000-4      	       3	 470402682 ns/op
BenchmarkUnionWith100-4     	  300000	      4333 ns/op
BenchmarkUnionWith1000-4    	  300000	      4359 ns/op
BenchmarkUnionWith10000-4   	  300000	      4403 ns/op
BenchmarkString100-4        	     200	   8369156 ns/op
BenchmarkString1000-4       	     200	   8389713 ns/op
BenchmarkString10000-4      	     200	   9655893 ns/op
PASS
ok  	github.com/negibokken/golang-training/ch11/ex07/intset32	30.843s
```

```
BenchmarkHas100-4           	  500000	      3700 ns/op
BenchmarkHas1000-4          	   50000	     37604 ns/op
BenchmarkHas10000-4         	    5000	    414067 ns/op
BenchmarkAdd100-4           	  200000	      9101 ns/op
BenchmarkAdd1000-4          	   20000	     88946 ns/op
BenchmarkAdd10000-4         	    2000	    886693 ns/op
BenchmarkAddAll100-4        	    5000	    236419 ns/op
BenchmarkAddAll1000-4       	      50	  37634570 ns/op
BenchmarkAddAll10000-4      	       1	3543767313 ns/op
BenchmarkUnionWith100-4     	   20000	     74789 ns/op
BenchmarkUnionWith1000-4    	    3000	    861121 ns/op
BenchmarkUnionWith10000-4   	     300	   8150325 ns/op
BenchmarkString100-4        	     100	  17630339 ns/op
BenchmarkString1000-4       	     100	  19700995 ns/op
BenchmarkString10000-4      	      50	  21698662 ns/op
PASS
ok  	github.com/negibokken/golang-training/ch11/ex07/mapintset	32.025s
```