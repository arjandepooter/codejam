import numpy as np
import algorithms as al

def generate_frequencies(runs):
    #first dimension is location, second is number put there
    freqs = np.zeros((1000,1000))
    for run in range(runs):
        x = al.rand_perm_2(1000)
        freqs[range(1000),x] += 1
    return freqs

def get_likelihood(x, freqs):
    return np.prod([float(1000-i)*freqs[i,x[i]]/sum(freqs[i, x[i:]]) for i in range(1000)])


def is_biased(x, freqs):
    if get_likelihood(x,freqs) > 1.0:
        return True
    else:
        return False

def get_output(infile, outfile, freqs):
    inputdata = open(infile, 'r')
    output = open(outfile, 'w')
    T = int(inputdata.readline())
    for t in range(T):
        N = int(inputdata.readline())
        x = [int(i) for i in inputdata.readline().split()]
        if not is_biased(x, freqs):
            quality = 'GOOD'
        else:
            quality ='BAD'
        output.write('Case #' + str(t+1) +': ' + quality + "\n")
    return None
