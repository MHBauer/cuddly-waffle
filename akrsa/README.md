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

# basic results 

hyperfine 'sbcl --non-interactive --load .\akrsa.lisp' 'python.exe .\akrsa.py'
Benchmark #1: sbcl --load .\akrsa.lisp
  Time (mean ± σ):      57.4 ms ±   2.5 ms    [User: 2.6 ms, System: 4.9 ms]
  Range (min … max):    52.4 ms …  67.9 ms    49 runs

  Warning: Statistical outliers were detected. Consider re-running this benchmark on a quiet PC without any interferences from other programs. It might help to use the '--warmup' or '--prepare' options.

Benchmark #2: python.exe .\akrsa.py
  Time (mean ± σ):     235.5 ms ±   2.7 ms    [User: 5.7 ms, System: 13.1 ms]
  Range (min … max):   231.7 ms … 240.6 ms    12 runs

Summary
  'sbcl --load .\akrsa.lisp' ran
    4.10 ± 0.18 times faster than 'python.exe .\akrsa.py'
