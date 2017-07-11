## スペック

```
% system_profiler SPHardwareDataType                                                  (git)-[master]
Hardware:

    Hardware Overview:

      Model Name: MacBook Pro
      Model Identifier: MacBookPro13,1
      Processor Name: Intel Core i5
      Processor Speed: 2 GHz
      Number of Processors: 1
      Total Number of Cores: 2
      L2 Cache (per Core): 256 KB
      L3 Cache: 4 MB
      Memory: 16 GB
```

## 実行結果

```
GOMAXPROCS=1 === RUN   Test_main
time: 1.30287853s
--- PASS: Test_main (1.30s)
PASS
ok  	_/Users/bokken/go/workspace/golang-training/ch09/ex06	1.378s
GOMAXPROCS=2
=== RUN   Test_main
time: 723.700567ms
--- PASS: Test_main (0.72s)
PASS
ok  	_/Users/bokken/go/workspace/golang-training/ch09/ex06	0.765s
GOMAXPROCS=3
=== RUN   Test_main
time: 560.393198ms
--- PASS: Test_main (0.56s)
PASS
ok  	_/Users/bokken/go/workspace/golang-training/ch09/ex06	0.595s
GOMAXPROCS=4
=== RUN   Test_main
time: 437.560285ms
--- PASS: Test_main (0.44s)
PASS
ok  	_/Users/bokken/go/workspace/golang-training/ch09/ex06	0.468s
```
