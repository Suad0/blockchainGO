package crypto

import "crypto/sha256"

type MerkleTreeNode struct {
	Left  *MerkleTreeNode
	Right *MerkleTreeNode
	Hash  []byte
}

// NewMerkleTree creates a Merkle tree from a list of data
func NewMerkleTree(data [][]byte) *MerkleTreeNode {
	var nodes []MerkleTreeNode

	for _, datum := range data {
		hash := sha256.Sum256(datum)
		nodes = append(nodes, MerkleTreeNode{Hash: hash[:]})
	}

	for len(nodes) > 1 {
		var newLevel []MerkleTreeNode
		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			var right MerkleTreeNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			} else {
				right = left
			}
			combinedHash := sha256.Sum256(append(left.Hash, right.Hash...))
			newLevel = append(newLevel, MerkleTreeNode{
				Left:  &left,
				Right: &right,
				Hash:  combinedHash[:],
			})
		}
		nodes = newLevel
	}
	return &nodes[0]
}
