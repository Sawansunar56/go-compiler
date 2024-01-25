package vm

import (
    "github.com/Sawansunar56/code"
    "github.com/Sawansunar56/compiler"
    "github.com/Sawansunar56/object"
)

const StackSize = 2048

type VM struct {
    constants []object.Object
    instructions code.Instructions

    stack []object.Object
    sp int // always points to the next value top of the stack is stack[sp-1]
}

func New(bytecode *compiler.Bytecode) *VM {
    return &VM {
        instructions: bytecode.Instructions,
        constants: bytecode.Constants,

        stack: make([]object.Object, StackSize),
        sp: 0,
    }
}

func (vm *VM) StackTop() object.Object {
    if vm.sp == 0 {
        return nil
    }
    return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
    for ip := 0; ip < len(vm.instructions); ip++ {
        op := code.Opcode(vm.instructions[ip])

        switch op {
        }
    }
    return nil
}
