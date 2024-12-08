package main

import (
	"strconv"
	"strings"
)

type vmContext struct {
	ip int
	instructions *[]instruction
	registers *[4]int
}

type instruction func(vm *vmContext) instruction

// AssemBunny interpreter
type AssemBunny struct {
	vm *vmContext
	instructions []instruction
	output func(value int) bool
}

func (asm AssemBunny) convert(value string) (bool, int) {
	if value[0] >= 'a' && value[0] <= 'z' {
		return true, int(value[0]-'a')
	}
	var v, _ = strconv.Atoi(value)
	return false, v
}

func (asm AssemBunny) nop1(reg bool, value int) instruction {
	return func(vm *vmContext) instruction { 
		if vm == nil { return asm.inc(reg, value) }
		vm.ip++ 
		return nil
	}
}

func (asm AssemBunny) nop2(reg1 bool, value1 int, reg2 bool, value2 int) instruction {
	return func(vm *vmContext) instruction { 
		if vm == nil { return asm.jnz(reg1, value1, reg2, value2 )}
		vm.ip++ 
		return nil
	}
}

func (asm AssemBunny) cpy(reg1 bool, value1 int, reg2 bool, value2 int) instruction {
	if (! reg2) { return asm.nop2(reg1, value1, reg2, value2) }
	if (reg1) {
		return func(vm *vmContext) instruction { 
			if vm == nil { return asm.jnz(reg1, value1, reg2, value2 )}
			vm.registers[value2] = vm.registers[value1]
			vm.ip++
			return nil
		}
	}

	return func(vm *vmContext) instruction { 
		if vm == nil { return asm.jnz(reg1, value1, reg2, value2 )}
		vm.registers[value2] = value1
		vm.ip++
		return nil
	}
}

func (asm AssemBunny) jnz(reg1 bool, value1 int, reg2 bool, value2 int) instruction {
	if (reg1 && reg2) {
		return func(vm *vmContext) instruction { 
			if vm == nil { return asm.cpy(reg1, value1, reg2, value2 )}
			if vm.registers[value1] != 0 {
				vm.ip += vm.registers[value2]
			} else {
				vm.ip++
			}
			return nil
		}
	} 
	if (reg1) {
		return func(vm *vmContext) instruction { 
			if vm == nil { return asm.cpy(reg1, value1, reg2, value2 )}
			if vm.registers[value1] != 0 {
				vm.ip += value2
			} else {
				vm.ip++
			}
			return nil
		}
	} 
	if value1 == 0 { return asm.nop2(reg1, value1, reg2, value2) }
	if reg2 {
		return func(vm *vmContext) instruction { 
			if vm == nil { return asm.cpy(reg1, value1, reg2, value2 )}
			vm.ip += vm.registers[value2] 
			return nil
		}
	}
	return func(vm *vmContext) instruction { 
		if vm == nil { return asm.cpy(reg1, value1, reg2, value2 )}
		vm.ip += value2 
		return nil
	}
}

func (asm AssemBunny) inc(reg bool, value int) instruction {
	if (! reg) { return asm.nop1(reg, value) }
	return func(vm *vmContext) instruction { 
		if vm == nil { return asm.dec(reg, value) }
		vm.registers[value]++ 
		vm.ip++;
		return nil
	}
}

func (asm AssemBunny) dec(reg bool, value int) instruction {
	if (! reg) { return asm.nop1(reg, value) }
	return func(vm *vmContext) instruction { 
		if vm == nil { return asm.inc(reg, value) }
		vm.registers[value]-- 
		vm.ip++;
		return nil
	}
}

func (asm AssemBunny) tgz(reg bool, value int) instruction {
	if (reg) { 
		return func(vm *vmContext) instruction {
			if vm == nil { return asm.inc(reg, value) }

			o := vm.ip + vm.registers[value]
			if o >= 0 && o < len(*vm.instructions) {
				i := (*vm.instructions)[o]
				i = i(nil) // generate toggled instruction
				(*vm.instructions)[o] = i
			}
			vm.ip++;
			return nil
		}
	}

	return func(vm *vmContext) instruction { 
		if vm == nil { return asm.inc(reg, value) }

		o := vm.ip + value
		if o >= 0 && o < len(*vm.instructions) {
			i := (*vm.instructions)[o]
			i = i(nil) // generate toggled instruction
			(*vm.instructions)[o] = i
		}
		vm.ip++;
		return nil
	}
}

func (asm AssemBunny) out(reg bool, value int) instruction {
	if (reg) {
		return func(vm *vmContext) instruction {
			if vm == nil { return asm.inc(reg, value) }
			
			if asm.output != nil { 
				if (asm.output(vm.registers[value])) {
					vm.ip = 1000; // terminate
				}
			}

			vm.ip++
			return nil
		}
	}

	return func(vm *vmContext) instruction {
		if vm == nil { return asm.inc(reg, value) }

		if asm.output != nil { 
			if (asm.output(value)) {
				vm.ip = 1000; // terminate
			}
		}

		vm.ip++
		return nil
	}
}

func (asm AssemBunny) parse(line string) instruction {
	switch line[0:3] {
		case "tgl": {
			reg, value := asm.convert(line[4:])
			return asm.tgz(reg, value)
		}

		case "cpy": {
			values := strings.Split(line[4:], " ")
			reg1, value1 := asm.convert(values[0])
			reg2, value2 := asm.convert(values[1])
			return asm.cpy(reg1, value1, reg2, value2);
		} 

		case "inc": {
			reg, value := asm.convert(line[4:])
			return asm.inc(reg, value)
		}

		case "dec": {
			reg, value := asm.convert(line[4:])
			return asm.dec(reg, value)
		}

		case "jnz": {
			values := strings.Split(line[4:], " ")
			reg1, value1 := asm.convert(values[0])
			reg2, value2 := asm.convert(values[1])
			return asm.jnz(reg1, value1, reg2, value2);
		}

		case "out": {
			reg, value := asm.convert(line[4:])
			return asm.out(reg, value)
		}

		default:
			return asm.nop1(false, 1)
	}
}

func (asm AssemBunny) load(day int) AssemBunny {
	instructions := []instruction {}

	for line := range readlines(day) {
		ins := asm.parse(line)
		instructions = append(instructions, ins)
	}

	asm.instructions = instructions
	return asm
}

func (asm AssemBunny) run(registers [4]int, preExecute func(ip int, registers *[4]int) int) [4]int {
	asm.vm = &vmContext {
		ip: 0,
		registers: &registers,
		instructions: &asm.instructions,
	}

	for asm.vm.ip >= 0 && asm.vm.ip < len(asm.instructions) {
		if preExecute != nil { asm.vm.ip = preExecute(asm.vm.ip, &registers) }
		if asm.vm.ip < 0 || asm.vm.ip >= len(asm.instructions) { break }

		ins := asm.instructions[asm.vm.ip]

		ins(asm.vm)
	}

	return registers
}