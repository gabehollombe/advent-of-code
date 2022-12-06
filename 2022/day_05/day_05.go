package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	crates string
}

type Instruction struct {
	amount    int
	fromIndex int
	toIndex   int
}

type State struct {
	stacks       []Stack
	instructions []Instruction
}

func loadInput(filename string) State {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(string(content), "\n")

	// split the lines into stack and instruction slices by finding the empty line between the two parts
	parts := strings.Split(strings.Join(lines, "\n"), "\n\n")
	stackLines := strings.Split(parts[0], "\n")
	instructionLines := strings.Split(parts[1], "\n")

	stacks := buildStacks(stackLines)
	instructions := buildInstructions(instructionLines)

	return State{stacks: stacks, instructions: instructions}
}

func buildInstructions(lines []string) []Instruction {
	// move <amount> from <moveFrom> to <moveTo>  (adjusts stack indexes to be 0-indexed)
	instructions := make([]Instruction, len(lines))
	for i, l := range lines {
		parts := strings.Split(l, " ")
		amount, _ := strconv.Atoi(parts[1])
		moveFrom, _ := strconv.Atoi(parts[3])
		moveTo, _ := strconv.Atoi(parts[5])
		instructions[i] = Instruction{amount: amount, fromIndex: moveFrom - 1, toIndex: moveTo - 1}
	}
	return instructions
}

func buildStacks(lines []string) []Stack {
	stackNumbers := strings.Split(lines[len(lines)-1], " ")
	numStacks, _ := strconv.ParseInt(stackNumbers[len(stackNumbers)-1], 10, 32)

	stacks := make([]Stack, numStacks)
	for _, l := range lines[:len(lines)-1] {
		for i := 0; i < int(numStacks); i++ {
			letterIndex := 2 + 4*i - 1 // where to pull the crate letter from in the line
			if len(l) < letterIndex {
				continue
			}
			letter := l[letterIndex]
			if letter != byte(' ') {
				// put letter into stack
				stacks[i].crates = string(letter) + stacks[i].crates
			}
		}
	}

	return stacks
}

func processAllInstructions(state State, moveAllAtOnce bool) State {
	var newStacks []Stack
	for _, inst := range state.instructions {
		if !moveAllAtOnce {
			newStacks = processInstructionOneByOne(state.stacks, inst)
			state = State{stacks: newStacks, instructions: state.instructions[1:]}
		} else {
			newStacks = processInstructionMoveAllAtOnce(state.stacks, inst)
			state = State{stacks: newStacks, instructions: state.instructions[1:]}
		}
	}
	return state
}

func processInstructionOneByOne(stacks []Stack, inst Instruction) []Stack {
	for i := 0; i < inst.amount; i++ {
		stacks = move(stacks, 1, inst.fromIndex, inst.toIndex)
	}
	return stacks
}

func processInstructionMoveAllAtOnce(stacks []Stack, inst Instruction) []Stack {
	stacks = move(stacks, inst.amount, inst.fromIndex, inst.toIndex)
	return stacks
}

// moves a single letter crate from <from> to <to>
func move(stacks []Stack, amount int, from int, to int) []Stack {
	fromStack := stacks[from]
	toStack := stacks[to]
	cratesToMove := fromStack.crates[len(fromStack.crates)-amount:]

	fromStack.crates = fromStack.crates[:len(fromStack.crates)-amount]
	toStack.crates += cratesToMove
	stacks[from] = fromStack
	stacks[to] = toStack
	return stacks
}

func main() {
	state := loadInput("input.txt")
	state = processAllInstructions(state, false)
	topCrates := ""
	for _, s := range state.stacks {
		topCrates += string(s.crates[len(s.crates)-1])
	}
	fmt.Printf("Part 1: %v\n", topCrates)

	state = loadInput("input.txt")
	state = processAllInstructions(state, true)
	topCrates = ""
	for _, s := range state.stacks {
		topCrates += string(s.crates[len(s.crates)-1])
	}
	fmt.Printf("Part 2: %v", topCrates)
}
