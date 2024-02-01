dep-graph:
	godepgraph -p github.com,google.golang.org,golang.org -s . | dot -Tpng > graph.png && open graph.png