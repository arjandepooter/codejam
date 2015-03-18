import numpy as np

def profile(string):
    letters = []
    letters.append(string[0])
    for a in string:
        if a != letters[-1]:
            letters.append(a)
    return letters

def counts(string):
    prof = profile(string)
    count = [0]*len(prof)
    position = 0
    for a in string:
        if a != prof[position]:
            position+=1
        count[position]+=1
    return count

def is_feasible(listofstrings):
    profs = map(profile, listofstrings)
    return all(profs[i]==profs[i+1] for i in range(len(profs)-1))

def minimum_moves(listofstrings):
    if is_feasible(listofstrings):
        count = np.array(map(counts, listofstrings))
        lengthofprofile = np.size(count, axis=1)
        moves = 0
        for i in range(lengthofprofile):
            smallest = min(count[:,i])
            largest = max(count[:,i])
            if smallest<largest:
                moves += min([sum(abs(count[:,i]-j)) for j in range(smallest,largest+1)])
        return str(moves)
    else:
        return "Fegla Won"

def get_output(instance):
    inputdata = open(instance + ".in", 'r')
    output = open(instance+ ".out", 'w')
    T = int(inputdata.readline())
    for t in range(T):
        N = int(inputdata.readline())
        stringlist =[]
        for n in range(N):
            stringlist.append(inputdata.readline())
        output.write('Case #' + str(t+1) +': ' + minimum_moves(stringlist) + "\n")
    return None
