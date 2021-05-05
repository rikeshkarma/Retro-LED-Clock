package placeholder

// for keeping things easy to read and type-safe
type Placeholder [5]string

// put the digits (placeholders) into variables
// using the placeholder array type above
var zero = Placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = Placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = Placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = Placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = Placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = Placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = Placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = Placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = Placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = Placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var Colon = Placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

// This array's type is "like": [10][5]string
//
// However:
// + "placeholder" is not equal to [5]string in type-wise.
// + Because: "placeholder" is a defined type, which is different
//   from [5]string type.
// + [5]string is an unnamed type.
// + placeholder is a named type.
// + The underlying type of [5]string and placeholder is the same:
//     [5]string
var Digits = [...]Placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}