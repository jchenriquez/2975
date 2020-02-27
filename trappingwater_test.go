package trappingwater

import (
  "testing"
  "os"
  "io"
  "bufio"
  "encoding/json"
)

type Test struct {
  Input []int `json:"input"`
  Output int `json:"output"`
}

func TestTrap (t *testing.T) {
  f, err := os.Open("tests.json")

  if err != nil {
    t.Error(err)
    return
  }

  defer f.Close()

  reader := bufio.NewReader(f)
  decoder := json.NewDecoder(reader)
  var tests map[string]Test

  for {
    err = decoder.Decode(&tests)

    if err == nil {
      for name, test := range tests {
        t.Run(name, func (k *testing.T) {
          result := Trap(test.Input)

          if result != test.Output {
            k.Errorf("result returned %d", result)
          }
        })
      }
    } else if err == io.EOF {
      break
    } else {
      t.Error(err)
      break
    }
  }
}