package length

import "fmt"

type Meter float64
type Inch float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (i Inch) String() string  { return fmt.Sprintf("%gin", i) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

func MToI(m Meter) Inch { return Inch(m * Meter(1/0.0254)) }
func MToF(m Meter) Feet { return Feet(m * Meter(1/0.3048)) }

func IToM(i Inch) Meter { return Meter(i * Inch(0.0254)) }
func IToF(i Inch) Feet  { return Feet(MToF(IToM(i))) }

func FToM(f Feet) Meter { return Meter(f * Feet(0.3048)) }
func FToI(f Feet) Inch  { return Inch(MToI(FToM(f))) }
