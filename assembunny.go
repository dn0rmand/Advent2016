package main

import (
	"strconv"
	"strings"
)

type vmContext struct {
	ip int
	registers *[4]int
}

type instruction func(vm *vmContext) 

// AssemBunny interpreter
type AssemBunny struct {
	vm *vmContext
	instructions []instruction
}

func (asm AssemBunny) convert(value string) (bool, int) {
	if value[0] >= 'a' && value[0] <= 'z' {
		return true, int(value[0]-'a')
	}
	var v, _ = strconv.Atoi(value)
	return false, v
}

func (asm AssemBunny) load(day int) AssemBunny {
	instructions := []instruction {}

	for line := range readlines(day) {
		var ins instruction

		switch line[0:3] {
			case "cpy": {
				_, r 		   := asm.convert(line[len(line)-1:])
				reg, value := asm.convert(line[4:len(line)-2])
				if reg {
					ins = func(vm *vmContext) { 
						vm.registers[r] = vm.registers[value] 
						vm.ip++
					}
				} else {
					ins = func(vm *vmContext) { 
						vm.registers[r] = value 
						vm.ip++
					}
				}
			} 
			case "inc": {
				_, r := asm.convert(line[4:])
				ins = func(vm *vmContext) { 
					vm.registers[r]++ 
					vm.ip++
				}
			}
			case "dec": {
				_, r := asm.convert(line[4:])
				ins = func(vm *vmContext) { 
					vm.registers[r]-- 
					vm.ip++
				}
			}
			case "jnz": {
				values := strings.Split(line[4:], " ")
				reg1, value1 := asm.convert(values[0])
				reg2, value2 := asm.convert(values[1])

				ins = func(vm *vmContext) { 
					if (reg1 && vm.registers[value1] == 0) || (!reg1 && value1 == 0) {
						vm.ip++;
					} else if reg2 { 
						vm.ip += vm.registers[value2]
					} else {
						vm.ip += value2
					}
				}
			}
		}

		instructions = append(instructions, ins)
	}

	asm.instructions = instructions
	return asm
}

func (asm AssemBunny) run(registers [4]int) [4]int {
	asm.vm = &vmContext {
		ip: 0,
		registers: &registers,
	}

	for asm.vm.ip < len(asm.instructions) {
		ins := asm.instructions[asm.vm.ip]

		ins(asm.vm)
	}

	return registers
}