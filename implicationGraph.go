/*
Implementation of implication graph.

An Implication Graph is a directed acyclic graph where each node represents a literal.
It consists on:
	* Node: assigned clause/literal.
	* Incident edge: represents the cause to that node assignment.

There are two types of nodes:
	* implied node -> this is forced during unit propagation.
	* decision node -> this was made by a guess. The decision node does not have any incident edge.

Each node (decision or implied) has a decision level associated with it.

Conflict: it happens when a variable x is set on true and false at the same time in the graph.
*/

package main

import "errors"

type implicationGraphNode struct {
	value         string
	state         bool
	impliedLevels []int
	impliedNodes  []*implicationGraphNode
}

type implicationGraph struct {
	decisionNodes map[string]*implicationGraphNode
	impliedNodes  map[string]*implicationGraphNode
	nodeByLevel   map[int][]*implicationGraphNode
}

func newimplicationGraph() (*implicationGraph, error) {
	return &implicationGraph{
		decisionNodes: make(map[string]*implicationGraphNode),
		impliedNodes:  make(map[string]*implicationGraphNode),
		nodeByLevel:   make(map[int][]*implicationGraphNode),
	}, nil
}

func (graph *implicationGraph) hasConflict(
	value string, state bool,
) (bool, *implicationGraphNode) {
	node, exists := graph.impliedNodes[value]
	if !exists {
		return false, nil
	}
	return node.state != state, node
}

func (graph *implicationGraph) addImpliedNode(
	value string, state bool, level int, parent *implicationGraphNode,
) (*implicationGraphNode, error) {
	isConflict, node := graph.hasConflict(value, state)
	if isConflict {
		//delete whole conflict level and send last decision node.
		return node, errors.New("GraphConflictError")
	}
	var impliedNode *implicationGraphNode
	if node == nil {
		impliedNode = &implicationGraphNode{
			value: value, state: state, impliedLevels: []int{level},
		}
		graph.impliedNodes[value] = impliedNode
	} else {
		impliedNode = graph.impliedNodes[value]
		impliedNode.impliedLevels = append(graph.impliedNodes[value].impliedLevels, level)
	}
	parent.impliedNodes = append(parent.impliedNodes, impliedNode)
	return graph.impliedNodes[value], nil
}

func (graph *implicationGraph) addDecisionNode(
	value string, state bool, level int,
) (*implicationGraphNode, error) {
	newNode := implicationGraphNode{
		value: value, state: state, impliedLevels: []int{level},
	}
	newNode.impliedLevels = append(newNode.impliedLevels, level)
	graph.decisionNodes[value] = &newNode
	return &newNode, nil
}
