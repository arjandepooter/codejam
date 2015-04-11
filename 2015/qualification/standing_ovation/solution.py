def peopleneeded(Smax, S):
    needed = 0
    for s in range(Smax+1):
        if sum(S[:s+1])<s+1:
            needed += s+1-sum(S[:s+1])
            S[s] += s+1-sum(S[:s+1])
    return needed

def get_output(instance):
    inputdata = open(instance + ".in", 'r')
    output = open(instance+ ".out", 'w')
    T = int(inputdata.readline())
    for t in range(T):
        Smax, S = inputdata.readline().split()
        Smax = int(Smax)
        S = [int(i) for i in list(S)]
        output.write('Case #' + str(t+1) +': ' + str(peopleneeded(Smax, S)) +  "\n")
    return None
