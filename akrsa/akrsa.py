import random
grb = random.getrandbits
e = grb(4096)
d = grb(4096)
m = grb(4096)
x = grb(4096)
c=(e*x)%m
r=(e*d)%m
y1=(r*x)%m
y2=(c*d)%m
print(y1)
print(y2)
print(y1==y2)
