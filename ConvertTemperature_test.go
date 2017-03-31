package tempconvert

import "testing"


func invokeFahrenheitToCelsius(input float64, expected_output float64, t *testing.T) {
     output := FahrenheitToCelsius(input)

     if output != expected_output {
          t.Errorf("FAILED: FarenheitToCelsius(%s) returned %d, expected %d", input, output, expected_output)
     } else {
          t.Logf("PASSED: The output was correctly %d", output)
     }
}

func invokeCelsiusToFahrenheit(input float64, expected_output float64, t *testing.T) {
     output := CelsiusToFahrenheit(input)

     if output != expected_output {
          t.Errorf("FAILED: CelsiusToFarenheit(%s) returned %d, expected %d", input, output, expected_output)
     } else {
          t.Logf("PASSED: The output was correctly %d", output)
     }
}

func invokeSimpleRound(input float64, expected_output float64, t *testing.T) {
     output := SimpleRound(input)

     if output != expected_output {
          t.Errorf("FAILED: CelsiusToFarenheit(%s) returned %d, expected %d", input, output, expected_output)
     } else {
          t.Logf("PASSED: The output was correctly %d", output)
     }
}

func TestFahrenheit32ToCelsius0(t *testing.T) {
     input := 32.0
     expected_output := 0.0

     invokeFahrenheitToCelsius(input, expected_output, t)
}

func TestFahrenheit50ToCelsius10(t *testing.T) {
     input := 50.0
     expected_output := 10.0

     invokeFahrenheitToCelsius(input, expected_output, t)
}

func TestFahrenheit40ToCelsius4dot44444(t *testing.T) {
     input := 40.0
     expected_output := 4.44444

     invokeFahrenheitToCelsius(input, expected_output, t)
}

func TestFahrenheitMinus40ToCelsiusMinus40(t *testing.T) {
     input := -40.0
     expected_output := -40.0

     invokeFahrenheitToCelsius(input, expected_output, t)
}

func TestFahrenheit69dot9998ToCelsius21dot111(t *testing.T) {
     input := 69.9998
     expected_output := 21.111

     invokeFahrenheitToCelsius(input, expected_output, t)
}


func TestCelsius32ToCelsius0(t *testing.T) {
     input := 0.0
     expected_output := 32.0

     invokeCelsiusToFahrenheit(input, expected_output, t)
}

func TestCelsius50ToFahrenheit10(t *testing.T) {
     input := 10.0
     expected_output := 50.0

     invokeCelsiusToFahrenheit(input, expected_output, t)
}

func TestCelsius40ToFahrenheit4dot44444(t *testing.T) {
     input := 4.44444
     expected_output := 39.99999

     invokeCelsiusToFahrenheit(input, expected_output, t)
}

func TestCelsiusMinus40ToFahrenheitMinus40(t *testing.T) {
     input := -40.0
     expected_output := -40.0

     invokeCelsiusToFahrenheit(input, expected_output, t)
}

func TestCelsius69dot9998ToFahrenheit21dot111(t *testing.T) {
     input := 21.111
     expected_output := 69.9998

     invokeCelsiusToFahrenheit(input, expected_output, t)
}
