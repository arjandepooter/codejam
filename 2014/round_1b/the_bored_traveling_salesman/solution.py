'''
graph = { "a" : ["c"],
          "b" : ["c", "e"],
          "c" : ["a", "b", "d", "e"],
          "d" : ["c"],
          "e" : ["c", "b"],
          "f" : []
        }
'''
import copy

def is_connected(graph):
    if graph=={}:
        return True
    seen = set([next(graph.iterkeys())])
    tocheck = seen.copy()
    while tocheck:
        node = tocheck.pop()
        tocheck = tocheck.union(set(graph[node]).difference(seen))
        seen = seen.union(graph[node])
    return len(seen)==len(graph)

def deletenodes(graph, nodes):
    for node in nodes:
        del graph[node]
    for i in graph.keys():
        for node in nodes:
            try:
                graph[i].remove(node)
            except:
                pass
    return graph

def bestroute(graph):
    tovisit = graph.keys()
    visited = [min(tovisit)]
    neighbors = set(graph[visited[0]])
    returnflight = list(visited)
    while len(visited)<len(tovisit):
        candidates = sorted(neighbors, reverse=True)
        candidateadded = False
        while not candidateadded:
            candidate = candidates.pop()
            testgraph = copy.deepcopy(graph)
            removedvertices = []
            node = returnflight.pop()
            while not candidate in graph[node]:
                removedvertices.append(node)
                node = returnflight.pop()
            returnflight.append(node)
            testgraph =deletenodes(testgraph, removedvertices)
            if is_connected(testgraph):
                graph=copy.deepcopy(testgraph)
                candidateadded=True
                visited.append(candidate)
                returnflight.append(candidate)
                neighbors=set()
                for node in returnflight:
                    neighbors = neighbors.union(set(graph[node]))
                    neighbors = neighbors.difference(visited)
            else:
                returnflight += reversed(removedvertices)
    return ''.join(visited)

def get_output(instance):
    inputdata = open(instance + ".in", 'r')
    output = open(instance+ ".out", 'w')
    T = int(inputdata.readline())
    for t in range(T):
        n,m = [int(i) for i in inputdata.readline().split()]
        graph = {}
        nodes = []
        for i in range(n):
            nodes.append(inputdata.readline()[:-1])
            graph[nodes[i]]=[]
        for i in range(m):
            edge = [int(i)-1 for i in inputdata.readline().split()]
            graph[nodes[edge[0]]].append(nodes[edge[1]])
            graph[nodes[edge[1]]].append(nodes[edge[0]])
        print t
        if t == 99:
            print graph
        output.write('Case #' + str(t+1) +': ' + bestroute(graph) +  "\n")
        print(len(bestroute(graph)))
    return None
