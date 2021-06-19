# The Prompt
The napkin sized piece of paper contains:

```
e=rand(4096)
d=rand(4096)
m=rand(4096)
r=(e*d)%m
e,m,r shared
x=rand(4096)
c=(x*e)%m
y=(c*d)%m
y=(x*r)%m
```

the equivalent construction is more apparent if some terms are moved around.
move e to the front, and move x to the rear.


```
c=(x*e)%m => c=(e*x)%m
y=(x*r)%m => y=(r*x)%m
```

r is a version of c

y has a leading term that is constructed of r or c, and a lagging term that is a random number.

# compile the necessary languages

```
sbcl --non-interactive --load .\akrsa-sbcl.lisp

go build -o akrsa-go .\akrsa.go

cd rs
cargo build --release
cp ./target/release/akrsa.exe ../akrsa-rs.exe
cd ../
```

# basic results 

## runtimes
windows 10 x64
```
hyperfine --warmup 5  'sbcl --load .\akrsa.lisp' 'sbcl --core akrsa-lisp' 'python.exe .\akrsa.py' '.\akrsa-go' '.\akrsa-rs'
Benchmark #1: sbcl --load .\akrsa.lisp
  Time (mean ± σ):      59.0 ms ±   2.2 ms    [User: 2.9 ms, System: 4.5 ms]
  Range (min … max):    53.7 ms …  63.6 ms    44 runs

Benchmark #2: sbcl --core akrsa-lisp
  Time (mean ± σ):      47.4 ms ±   3.3 ms    [User: 1.9 ms, System: 7.2 ms]
  Range (min … max):    41.8 ms …  59.4 ms    53 runs

Benchmark #3: python.exe .\akrsa.py
  Time (mean ± σ):     234.1 ms ±   1.7 ms    [User: 5.7 ms, System: 11.2 ms]
  Range (min … max):   231.7 ms … 237.5 ms    12 runs

Benchmark #4: .\akrsa-go
  Time (mean ± σ):      11.4 ms ±   0.4 ms    [User: 1.8 ms, System: 4.2 ms]
  Range (min … max):     9.4 ms …  13.4 ms    159 runs

Benchmark #5: .\akrsa-rs
  Time (mean ± σ):       6.9 ms ±   0.4 ms    [User: 1.9 ms, System: 4.2 ms]
  Range (min … max):     5.5 ms …   8.7 ms    212 runs

Summary
  '.\akrsa-rs' ran
    1.66 ± 0.11 times faster than '.\akrsa-go'
    6.91 ± 0.60 times faster than 'sbcl --core akrsa-lisp'
    8.59 ± 0.57 times faster than 'sbcl --load .\akrsa.lisp'
   34.08 ± 1.86 times faster than 'python.exe .\akrsa.py'
```

## compiled sizes

```
  2324480 akrsa-go.exe
 40639072 akrsa-lisp
   224768 akrsa-rs.exe
```
