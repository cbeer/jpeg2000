package jpeg2000

import (
	"image"
	"os"
  "testing"
)

func TestHello(t *testing.T) {
  testCases := []string{
	   "./testdata/jp2",
   }

   for _, tc := range testCases {
     _, err := decodeFile(tc)
     if err != nil {
    		t.Errorf("%s: %v", tc+".jpeg", err)
    		continue
     }
   }
}

func decodeFile(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Decode(f)
}
