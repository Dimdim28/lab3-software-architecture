package test

import (
	"image/color"
	"strings"
	"testing"

	"github.com/Dimdim28/lab3-software-architecture/painter"
	"github.com/Dimdim28/lab3-software-architecture/painter/lang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parse_struct(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.Operation
	}{ // test cases, for the checking loop =)
		{
			name:    "background rectangle",
			command: "bgrect 0 0 100 100",
			op:      &painter.BgRectangle{X1: 0, Y1: 0, X2: 100, Y2: 100},
		},
		{
			name:    "figure",
			command: "figure 200 200",
			op:      &painter.Figure{X: 200, Y: 200, C: color.RGBA{R: 255, G: 255, B: 0, A: 1}},
		},
		{
			name:    "move",
			command: "move 100 100",
			op:      &painter.Move{X: 100, Y: 100},
		},
		{
			name:    "update",
			command: "update",
			op:      painter.UpdateOp,
		},
		{
			name:    "invalid command",
			command: "invalidcommand",
			op:      nil,
		},
	}
	// created parser
	// parser := &lang.Parser{}
	//ya poshel spat, zavtra utrom dopishu eshe commentov esli ne len budet.

	for _, tc := range tests { //created checking loop
		t.Run(tc.name, func(t *testing.T) {
			parser := &lang.Parser{}
			ops, err := parser.Parse(strings.NewReader(tc.command)) //parse the command string and get the list of operations
			if tc.op == nil {                                       //Check if the expected operation is nil
				assert.Error(t, err) // then assert that an error was returned. Vanya, it is comment for Oleg
			} else { //  when he will return, he will not waste too much time to understand our govnocode
				require.NoError(t, err)         // checking errors
				//require.Len(t, ops, 1)          // Check that the length of the resulting operations slice is 1. Типа онли по 1 команде с каждой строки парсится.
				assert.IsType(t, tc.op, ops[1]) // сheck that the type of the parsed operation matches the expected type.
				assert.Equal(t, tc.op, ops[1])  //  алерт иф зей аре нот икуал
			}
		})
	}
}

func Test_parse_func(t *testing.T) {
	tests := []struct { // test cases, что же это ещё может быть
		name    string
		command string
		op      painter.Operation
	}{
		{
			name:    "white fill",
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill),
		},
		{
			name:    "green fill",
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill),
		},
		{
			name:    "reset screen",
			command: "reset",
			op:      painter.OperationFunc(painter.ResetScreen),
		},
	}
	// created Parser object
	parser := &lang.Parser{}
	// checking loop =)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := parser.Parse(strings.NewReader(tc.command))

			
			// выше я уже коментил эту шнягу, но могу и тут, я уже выспался, могу и написать.
			require.NoError(t, err)         //checking for no errors =)    (да кто бы мог подумать)
			require.Len(t, ops, 1)          // Check that the length of the resulting operations slice is 1. Expect only one operation to be parsed from each command string.
			assert.IsType(t, tc.op, ops[0]) // Check that the type of the parsed operation matches the expected type.

		})
	}
}
