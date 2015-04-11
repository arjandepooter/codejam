import numpy as np
from collections import Counter
N = 7
edges = {(4,5), (4,2), (1,2), (3,1), (6,4), (3,7)}
count = np.zeros(N)
for i in edges:
    for j in i:
        count[j-1] +=1

def is_binary_tree(edgeset):
    found_root = False
    count = np.zeros(N)
    for i in edges:
        for j in i:
            count[j] +=1
    count = Counter(count)
    for node, occurences in count.items():
        if occurences == 2 and not found_root:
            found_root = True
        elif occurences == 2:
            return False
        elif occurences != 1 and occurences != 3:
            return False
    return True
