package freezer

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	f := New("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiI1NGU3YzkzZC1kN2I4LTQ2NjYtOGQ0Yi1mNzQzY2M4MGNkNTciLCJlbWFpbCI6ImNhaWNoYW8ueHVAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siaWQiOiJOWUMxIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6IjU1MmE4M2IxZmJmZjQzYWY1MzZhIiwic2NvcGVkS2V5U2VjcmV0IjoiNjYxOTg4MTRjYTZjZmJjNzZlMDZlZDEyZmM5ZWQxYzEwZDM1MjRiZWNhMjc0YTA2OTE2Zjk3NTE1ZTM5OTBlMiIsImlhdCI6MTY3MDg0Nzk5NH0.C8TNhbFV-GzkExEuPYJxP-x8bxo6VBInFMdMUT4B-5Y")
	file, err := os.Open("./testdata/ipfs.json")
	require.Nil(t, err)

	cid, err := f.PinFile(file)
	require.Nil(t, err)

	fmt.Println(cid)

	cid2, err := f.PinJson(map[string]interface{}{
		"name":        "Azuki",
		"description": "Azuki....",
	})
	require.Nil(t, err)

	fmt.Println(cid2)

	cid3, err := f.PinERC1155(map[string]interface{}{
		"name":        "Azuki",
		"description": "Azuki....",
	}, file)
	require.Nil(t, err)

	fmt.Println(cid3)
}
