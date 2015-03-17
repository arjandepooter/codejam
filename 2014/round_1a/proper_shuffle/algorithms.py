import random
import numpy as np

def rand_perm_1(N):
    a = range(N)
    for k in range(N):
        p = random.randint(k,N-1)
        swap=a[p]
        a[p]=a[k]
        a[k]=swap
    return a

def rand_perm_2(N):
    a = range(N)
    for k in range(N):
        p = random.randint(0,N-1)
        swap=a[p]
        a[p]=a[k]
        a[k]=swap
    return a

def size_intersection(N):
    x = []
    y = []
    for i in range(120):
        x.append(np.std(rand_perm_2(1000)[:N]))
        y.append(np.std(rand_perm_1(1000)[:N]))
    x.sort()
    y.sort()
    overlapy = sum(i<x[119] for i in y)
    overlapx = sum(i>y[0] for i in x)
    return overlapy+overlapx
