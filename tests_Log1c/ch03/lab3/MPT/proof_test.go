package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/trie"
	"github.com/stretchr/testify/require"
)


func address_and_value(size int) ([][]byte,  []byte){
	random := rand.New(rand.NewSource(0))

	keys := make([][]byte, size)
	for i := 0; i < size; i++ {
		k := make([]byte, 20)
		random.Read(k)
		keys[i] = k
	}
	value := make([]byte, 100)
	random.Read(value)
	return keys, value
}



func TestEthProof_size(t *testing.T) {
	mpt := new(trie.Trie)
	size := 2000
	keys, val := address_and_value(size)
	for i := 0; i < size; i++ {
		mpt.Update(keys[i], val)
	}
	w := NewProofDB()
	_ = mpt.Prove(keys[0], 0, w)
	//fmt.Println(w.kv)
	fmt.Println(len(w.kv))
	sum := 0
	for _, v := range w.kv {
		//fmt.Println(k)
		fmt.Println(len(v))
		curLen := len(v) / 32
		fmt.Println(curLen - 1)
		sum += curLen - 1
	}
	fmt.Println(sum)
}




func TestEthProof(t *testing.T) {
	mpt := new(trie.Trie)
	mpt.Update([]byte{1, 2, 3}, []byte("hello"))
	mpt.Update([]byte{1, 2, 3, 4, 5}, []byte("world"))
	w := NewProofDB()
	err := mpt.Prove([]byte{1, 2, 3}, 0, w)
	require.NoError(t, err)
	rootHash := mpt.Hash()
	val, err := trie.VerifyProof(rootHash, []byte{1, 2, 3}, w)
	require.NoError(t, err)
	require.Equal(t, []byte("hello"), val)
	fmt.Printf("root hash: %x\n", rootHash)
}

func TestMyTrie(t *testing.T) {
	tr := NewTrie()
	tr.Put([]byte{1, 2, 3}, []byte("hello"))
	tr.Put([]byte{1, 2, 3, 4, 5}, []byte("world"))
	n0, ok := tr.root.(*ExtensionNode)
	require.True(t, ok)
	n1, ok := n0.Next.(*BranchNode)
	require.True(t, ok)
	fmt.Printf("n0 hash: %x, Serialized: %x\n", n0.Hash(), n0.Serialize())
	fmt.Printf("n1 hash: %x, Serialized: %x\n", n1.Hash(), n1.Serialize())
}

func TestProveAndVerifyProof(t *testing.T) {
	t.Run("should not generate proof for non-exist key", func(t *testing.T) {
		tr := NewTrie()
		tr.Put([]byte{1, 2, 3}, []byte("hello"))
		tr.Put([]byte{1, 2, 3, 4, 5}, []byte("world"))
		notExistKey := []byte{1, 2, 3, 4}
		_, ok := tr.Prove(notExistKey)
		require.False(t, ok)
	})

	t.Run("should generate a proof for an existing key, the proof can be verified with the merkle root hash", func(t *testing.T) {
		tr := NewTrie()
		tr.Put([]byte{1, 2, 3}, []byte("hello"))
		tr.Put([]byte{1, 2, 3, 4, 5}, []byte("world"))

		key := []byte{1, 2, 3}
		proof, ok := tr.Prove(key)
		require.True(t, ok)

		rootHash := tr.Hash()

		// verify the proof with the root hash, the key in question and its proof
		val, err := VerifyProof(rootHash, key, proof)
		require.NoError(t, err)

		// when the verification has passed, it should return the correct value for the key
		require.Equal(t, []byte("hello"), val)
	})

	t.Run("should fail the verification of the trie was updated", func(t *testing.T) {
		tr := NewTrie()
		tr.Put([]byte{1, 2, 3}, []byte("hello"))
		tr.Put([]byte{1, 2, 3, 4, 5}, []byte("world"))

		// the hash was taken before the trie was updated
		rootHash := tr.Hash()

		// the proof was generated after the trie was updated
		tr.Put([]byte{5, 6, 7}, []byte("trie"))
		key := []byte{1, 2, 3}
		proof, ok := tr.Prove(key)
		require.True(t, ok)

		// should fail the verification since the merkle root hash doesn't match
		_, err := VerifyProof(rootHash, key, proof)
		require.Error(t, err)
	})
}
