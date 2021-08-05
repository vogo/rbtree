/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rbtree

const (
	maxStackDeep = 32
)

type stack struct {
	nodes     []*Node
	positions []Position
	index     int
}

func newStack(root *Node) *stack {
	s := &stack{
		nodes:     make([]*Node, maxStackDeep),
		positions: make([]Position, maxStackDeep),
		index:     -1,
	}
	s.push(&Node{}, Left)
	s.nodes[0].Left = root
	return s
}

func (s *stack) init(root *Node) {
	s.nodes[0].Left = root
}

func (s *stack) reset() {
	for s.index > 0 {
		s.nodes[s.index] = nil
		s.index--
	}

	s.nodes[0].Left = nil
}

func (s *stack) push(n *Node, pos Position) {
	s.index++
	s.nodes[s.index], s.positions[s.index] = n, pos
}

func (s *stack) insertBefore(n *Node, pos Position) {
	s.nodes[s.index+1], s.positions[s.index+1] = s.nodes[s.index], s.positions[s.index]
	s.nodes[s.index], s.positions[s.index] = n, pos
	s.index++
}

func (s *stack) pop() *stack {
	if s.index == 0 {
		return s
	}

	s.nodes[s.index] = nil
	s.index--
	return s
}

func (s *stack) root() *Node {
	return s.nodes[0].Left
}

func (s *stack) node() *Node {
	return s.nodes[s.index]
}

func (s *stack) position() Position {
	return s.positions[s.index]
}

func (s *stack) parentPosition() Position {
	return s.positions[s.index-1]
}

func (s *stack) parent() *Node {
	if s.index > 0 {
		return s.nodes[s.index-1]
	}
	return nil
}

func (s *stack) sibling() *Node {
	if s.index > 0 {
		i := s.index - 1
		if s.positions[i] == Left {
			return s.nodes[i].Right
		} else {
			return s.nodes[i].Left
		}
	}
	return nil
}

func (s *stack) childSibling() *Node {
	if s.position() == Left {
		return s.node().Right
	} else {
		return s.node().Left
	}
}

func (s *stack) bindChild(n *Node) {
	if s.position() == Left {
		s.node().Left = n
	} else {
		s.node().Right = n
	}
}
