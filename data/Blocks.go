package data

import (
	"encoding/json"
	"log"
	"os"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

type Blockchain struct {
	Data []Block //Blockchain
	Info string  //Information ?
}

var FileName string = "data/Blockchain.json"

// append blockchain
func (chain *Blockchain) AppendBlock(b Block) {
	chain.Data = append(chain.Data, b)
}

// Save entire blockchain,
func (b *Blockchain) Save(File string) {
	f, err := os.Create(File)
	if err != nil {
		log.Fatal("Create file wrong: ", err)
	}
	defer f.Close()

	content, err := json.Marshal(b)
	if err != nil {
		log.Fatal("Marshalling Blockchain Wrong: ", err)
	}
	_, err = f.Write(content)
	if err != nil {
		log.Fatal("Writing Blockchian in JSON Wronh: ", err)
	}

}

// Reload Blockchain
func Load_Blockchain(File string) Blockchain {
	f, err := os.ReadFile(File)
	if err != nil {
		log.Fatal("Error reading data.json: ", err)
	}

	var filecontent Blockchain // read in blockchain
	err = json.Unmarshal(f, &filecontent)
	if err != nil {
		log.Fatal("Error Unmarshalling data.json into Blockchain: ", err)
	}

	return filecontent
}
