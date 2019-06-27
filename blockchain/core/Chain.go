package core

type Chain struct {
	Blocks []Block
}

func GenerateChain() *Chain {
	chain := Chain{}
	chain.AppendBlock(GenerateGenesisBlock())
	return &chain
}

func (this *Chain) AppendBlock(newBlock Block) {
	if (len(this.Blocks) == 0) {
		this.Blocks = append(this.Blocks, newBlock)
		return
	}

	if newBlock.IsValid(this.Blocks[len(this.Blocks)-1]) {
		this.Blocks = append(this.Blocks, newBlock)
	} else {
		println("invalid block")
	}
}

func (this *Chain) SendData(data string) {
	newBlock := GenerateBlock(this.Blocks[len(this.Blocks)-1], data)
	this.AppendBlock(newBlock)
}
