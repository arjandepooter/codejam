import algorithms
from scipy import stats

def find_best_alpha(N, nrtests=10):
    #Returns the best alpha to test H_0: the generator is unbiased if we look for elements that ended up less than N spaces from their initial location
    bestalpha=0
    besterrors = 120
    for alpha in [0.80, 0.85, 0.90,0.95]:
        errors = (1-alpha)*60
        threshold = stats.binom.ppf(alpha,p=N/1000.0, n=1000-N)
        for tests in range(nrtests*60):
            x = algorithms.rand_perm_2(1000)
            if sum(i < x[i] <= i+N for i in range(1000-N))<threshold:
                errors += 1.0/nrtests
        if errors < besterrors:
            besterrors = errors
            bestalpha = alpha
    return bestalpha, besterrors

def is_unbiased(x, N=250, alpha=0.9):
    threshold = stats.binom.ppf(alpha,p=N/1000.0, n=1000-N)
    if sum(i < x[i] <= i+N for i in range(1000-N))<threshold:
        return True
    else:
        return False

def get_output(infile, outfile, band=250):
    inputdata = open(infile, 'r')
    output = open(outfile, 'w')
    T = int(inputdata.readline())
    for t in range(T):
        N = int(inputdata.readline())
        x = [int(i) for i in inputdata.readline().split()]
        if is_unbiased(x, N=band):
            quality = 'GOOD'
        else:
            quality ='BAD'
        output.write('Case #' + str(t+1) +': ' + quality + "\n")
    return None
