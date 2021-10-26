package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	test "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/testcommon"
	"github.com/awalterschulze/gographviz"
)

func main() {

	/*
		1 lvl of async calls
	*/
	// callGraph := test.CreateGraphTestOneAsyncCall()
	// callGraph := test.CreateGraphTestOneAsyncCallNoCallback()
	// callGraph := test.CreateGraphTestOneAsyncCallFail()
	// callGraph := test.CreateGraphTestOneAsyncCallNoCallbackFail()
	// callGraph := test.CreateGraphTestAsyncCallIndirectFail()
	// callGraph := test.CreateGraphTestOneAsyncCallbackFail()
	// callGraph := test.CreateGraphTestAsyncCallbackIndirectFail()
	// callGraph := test.CreateGraphTestAsyncCallIndirectFailCrossShard()
	// callGraph := test.CreateGraphTestOneAsyncCallFailCrossShard()
	// callGraph := test.CreateGraphTestOneAsyncCallbackFailCrossShard()
	// callGraph := test.CreateGraphTestTwoAsyncCallsSecondCallbackFailLocalCross()
	// callGraph := test.CreateGraphTestAsyncCallbackIndirectFailCrossShard()
	// callGraph := test.CreateGraphTestSyncCalls()
	// callGraph := test.CreateGraphTestSyncCalls2()
	// callGraph := test.CreateGraphTestOneAsyncCall()
	// callGraph := test.CreateGraphTestOneAsyncCallCrossShard()
	// callGraph := test.CreateGraphTestOneAsyncCallCrossShard2() //!
	// callGraph := test.CreateGraphTestOneAsyncCallNoCallbackCrossShard()
	// callGraph := test.CreateGraphTestOneAsyncCallFailNoCallbackCrossShard()
	// callGraph := test.CreateGraphTestTwoAsyncCalls()
	// callGraph := test.CreateGraphTestTwoAsyncCallsOneFail()
	// callGraph := test.CreateGraphTestTwoAsyncCallsLocalCross()
	// callGraph := test.CreateGraphTestTwoAsyncCallsCrossLocal()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncSecondFail()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncLocalCross()
	// callGraph := test.CreateGraphTestCallbackCallsSync()
	// callGraph := test.CreateGraphTestSyncAndAsync1()
	// callGraph := test.CreateGraphTestSyncAndAsync2()
	// callGraph := test.CreateGraphTestSyncAndAsync3()
	// callGraph := test.CreateGraphTestSyncAndAsync6()
	// callGraph := test.CreateGraphTestSyncAndAsync7()
	// callGraph := test.CreateGraphTestSyncAndAsync8()
	// callGraph := test.CreateGraphTestTwoAsyncCallsCrossShard()
	// callGraph := test.CreateGraphTestTwoAsyncCallsFirstCallbackFailCrossShard()
	// callGraph := test.CreateGraphTestSyncCallsFailPropagation()
	// callGraph := test.CreateGraphTestTwoAsyncCallsFirstFail()
	// callGraph := test.CreateGraphTestTwoAsyncCallsFirstFailLocalCross()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncFirstNoCallbackLocalCross()
	callGraph := test.CreateGraphTestOneAsyncCallCustomGasLocked()

	/*
		multi lvl of async calls
	*/
	// callGraph := test.CreateGraphTestAsyncCallsAsync()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncLocalCross()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncCrossShard()
	// callGraph := test.CreateGraphTestAsyncsOnMultiLevelFail1()
	// callGraph := test.CreateGraphTestCallbackCallsAsyncCrossCross()
	// callGraph := test.CreateGraphTestAsyncCallsCrossShard6()
	// callGraph := test.CreateGraphTestAsyncCallsCrossShard7()
	// callGraph := test.CreateGraphTestSyncAndAsync5()
	// callGraph := test.CreateGraphTestDifferentTypeOfCallsToSameFunction()
	// callGraph := test.CreateGraphTestCallbackCallsAsyncLocalLocal()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncSecondFail()
	// callGraph := test.CreateGraphTestCallbackCallsAsyncCrossLocal()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncSecondCallbackFailCrossShard()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncBothCallbacksFailLocalCross()
	// callGraph := test.CreateGraphTestAsyncCallsAsyncSecondCallbackFailLocalCross()

	///////////////////

	graphviz := toGraphviz(callGraph, true)
	createSvg("1 call-graph", graphviz)

	executionGraph := callGraph.CreateExecutionGraphFromCallGraph()
	graphviz = toGraphviz(executionGraph, true)
	createSvg("2 execution-graph", graphviz)

	gasGraph := executionGraph.ComputeGasGraphFromExecutionGraph()
	gasGraph.PropagateSyncFailures()
	gasGraph.AssignExecutionRounds(nil)

	graphviz = toGraphviz(gasGraph, false)
	createSvg("3 initial-gas-graph", graphviz)

	gasGraph.ComputeRemainingGasBeforeCallbacks(nil)
	graphviz = toGraphviz(gasGraph, false)
	createSvg("4 gas-graph-gasbeforecallbacks", graphviz)

	gasGraph.ComputeRemainingGasAfterCallbacks()
	graphviz = toGraphviz(gasGraph, false)
	createSvg("5 gas-graph-gasaftercallbacks-norestore", graphviz)
}

func createSvg(file string, graphviz *gographviz.Graph) {
	location := os.Args[1]

	destDot := location + file + ".dot"

	output := graphviz.String()
	err := ioutil.WriteFile(destDot, []byte(output), 0644)
	if err != nil {
		panic(err)
	}

	out, err := exec.Command("dot" /*"-extent 800x1500",*/, "-Tsvg", destDot).Output()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(location+file+".svg", out, 0644)
	if err != nil {
		panic(err)
	}
}

func toGraphviz(graph *test.TestCallGraph, showGasEdgeLabels bool) *gographviz.Graph {
	graphviz := gographviz.NewGraph()
	graphviz.Directed = true
	graphName := "G"
	graphviz.Attrs["nodesep"] = "1.5"

	nodeCounters := make(map[string]int)
	for _, node := range graph.Nodes {
		node.Label, node.VisualLabel = computeUniqueGraphvizNodeLabel(node, nodeCounters)
	}

	for _, node := range graph.Nodes {
		nodeAttrs := make(map[string]string)
		setNodeAttributes(node, nodeAttrs)
		from := node.Label
		graphviz.AddNode(graphName, from, nodeAttrs)
		for _, edge := range node.GetEdges() {
			to := edge.To.Label
			edgeAttrs := make(map[string]string)
			if edge.Label != "" {
				setEdgeLabel(edgeAttrs, edge, showGasEdgeLabels)
			}
			setEdgeAttributes(edge, edgeAttrs)
			graphviz.AddEdge(from, to, true, edgeAttrs)
		}
	}

	return graphviz
}

func setNodeAttributes(node *test.TestCallNode, attrs map[string]string) {
	if node.IsStartNode {
		attrs["shape"] = "box"
	}
	// if node.Visited {
	// 	attrs["penwidth"] = "4"
	// }
	setGasLabelForNode(node, attrs)
	if !node.IsGasLeaf() {
		if node.Fail || node.IsIncomingEdgeFail() || node.HasFailSyncEdge() {
			attrs["fillcolor"] = "hotpink"
		} else {
			attrs["fillcolor"] = "lightgrey"
		}
		attrs["style"] = "filled"
		attrs["label"] = node.VisualLabel
	}
}

func setEdgeLabel(attrs map[string]string, edge *test.TestCallEdge, showGasEdgeLabels bool) {
	attrs["label"] = edge.Label
	if showGasEdgeLabels && edge.Type != test.Callback && edge.Type != test.CallbackCrossShard {
		attrs["label"] += "\n" +
			"P" + strconv.Itoa(int(edge.GasLimit)) +
			"/U" + strconv.Itoa(int(edge.GasUsed))
		if edge.Type == test.Async || edge.Type == test.AsyncCrossShard {
			attrs["label"] += "/CU" + strconv.Itoa(int(edge.GasUsedByCallback))
		}
	}
	attrs["label"] = strconv.Quote(attrs["label"])
}

func setEdgeAttributes(edge *test.TestCallEdge, attrs map[string]string) {
	if edge.To.IsGasLeaf() {
		attrs["color"] = "black"
		return
	}
	switch edge.Type {
	case test.Sync:
		attrs["color"] = "blue"
	case test.Async:
		attrs["color"] = "red"
	case test.AsyncCrossShard:
		attrs["color"] = "red"
		attrs["style"] = "dashed"
	case test.Callback:
		attrs["color"] = "grey"
	case test.CallbackCrossShard:
		attrs["color"] = "grey"
		attrs["style"] = "dashed"
	default:
		attrs["color"] = "black"
	}
}

// generates unique graphviz node label using "_number" suffix if necessary
func computeUniqueGraphvizNodeLabel(node *test.TestCallNode, nodeCounters map[string]int) (string, string) {
	if nodeCounters == nil {
		return node.Label, node.Label
	}
	if node.VisualLabel != "" {
		return node.Label, node.VisualLabel
	}

	var prefix string
	if node.Call.FunctionName == test.LeafLabel {
		prefix = test.LeafLabel
	} else {
		prefix, _ = strconv.Unquote(node.Label)
	}

	counter, present := nodeCounters[prefix]
	if !present {
		counter = 0
	}
	counter++
	nodeCounters[prefix] = counter

	suffix := ""
	if counter > 1 {
		suffix = "_" + strconv.Itoa(counter)
	}

	var visualLabel string
	if node.WillNotExecute() {
		visualLabel = strconv.Quote(prefix)
	} else {
		visualLabel = strconv.Quote(fmt.Sprintf("%s [%d]", prefix, node.ExecutionRound))
	}

	return strconv.Quote(prefix + suffix), visualLabel
}

const gasFontStart = "<<font color='green'>"
const gasFontEnd = "</font>>"

func setGasLabelForNode(node *test.TestCallNode, attrs map[string]string) {
	if node.GasLimit == 0 && node.GasUsed == 0 {
		// special label for end nodes without gas info
		if node.IsGasLeaf() {
			attrs["label"] = strconv.Quote("*")
		}
		return
	}

	gasLimit := strconv.Itoa(int(node.GasLimit))
	gasUsed := strconv.Itoa(int(node.GasUsed))
	gasRemaining := strconv.Itoa(int(node.GasRemaining))
	gasAccumulated := strconv.Itoa(int(node.GasAccumulated))
	gasLocked := strconv.Itoa(int(node.GasLocked))
	var xlabel string
	if node.IsGasLeaf() {
		if node.WillNotExecute() {
			attrs["label"] = strconv.Quote(test.LeafLabel)
		} else {
			parent := node.Parent
			if node.IsGasLeaf() && parent != nil && parent.IsIncomingEdgeFail() {
				attrs["label"] = strconv.Quote(test.LeafLabel)
			} else {
				attrs["label"] = gasFontStart + gasUsed + gasFontEnd
			}
		}
	} else {
		// display only gas locked for uncomputed gas values (for group callbacks and context callbacks)
		if node.GasLimit == 0 || node.WillNotExecute() {
			return
		}
		xlabel = gasFontStart
		xlabel += "P" + gasLimit
		if node.GasLocked != 0 {
			xlabel += "/L" + gasLocked
		}

		xlabel += "<br/>R" + gasRemaining
		if node.GasAccumulated != 0 {
			xlabel += "<br/>A" + gasAccumulated
		}
		xlabel += gasFontEnd
		attrs["xlabel"] = xlabel
	}
}
