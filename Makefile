dep-graph:
	godepgraph -o github.com/simp7/auth-server github.com/simp7/auth-server | dot -Tpng > graph.png && open graph.png