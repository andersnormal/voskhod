package registry

// import (
// 	pb "github.com/andersnormal/voskhod/proto"
// )

// func addAgent(old, neu []*registry.Node) []*registry.Node {
// 	for _, n := range neu {
// 		var seen bool
// 		for i, o := range old {
// 			if o.Id == n.Id {
// 				seen = true
// 				old[i] = n
// 				break
// 			}
// 		}
// 		if !seen {
// 			old = append(old, n)
// 		}
// 	}
// 	return old
// }

// func addAgents(old, neu []*pb.Agent) []*pb.Agent {
// 	for _, s := range neu {
// 		var seen bool
// 		for i, o := range old {
// 			if o.Version == s.Version {
// 				s.Nodes = addNodes(o.Nodes, s.Nodes)
// 				seen = true
// 				old[i] = s
// 				break
// 			}
// 		}
// 		if !seen {
// 			old = append(old, s)
// 		}
// 	}
// 	return old
// }
