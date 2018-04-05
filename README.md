[![Godoc](https://img.shields.io/badge/Godoc-mnemonic-blue.svg)](https://godoc.org/github.com/CasalettoJ/mnemonic/mnemonic)


mnemonic generates a 24 mnemonic sentence and 512bit seed following

https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki

https://github.com/bitcoinbook/bitcoinbook/blob/develop/ch05.asciidoc

## Notes

run `dep ensure`

## Usage

Documentation:
https://godoc.org/github.com/CasalettoJ/mnemonic/mnemonic




### Sample output from running `RunDemo()`:

```
Step 1: Create 256 Random bits: 904fb2763eba9b0d7e4b83e651bda6bd632eb9598bf7bb6bd168f4be94b884df

Step 2: Generate a checksum:
9209c784a9143265184edea47c120b946f775b155124fbc98175148072235b22
and take the first (bits/32) bits:
92

Step 3: Add the checksum to the end of the generated bits:
904fb2763eba9b0d7e4b83e651bda6bd632eb9598bf7bb6bd168f4be94b884df92

Step 4: split the result into 11-bit-length segments
10010000010 | 01111101100 | 10011101100 | 01111101011 | 10101001101 | 10000110101 | 11111001001 | 01110000011 | 11100110010 | 10001101111 | 01101001101 | 01111010110 | 00110010111 | 0101110
0101 | 01100110001 | 01111110111 | 10111011011 | 01011110100 | 01011010001 | 11101001011 | 11101001010 | 01011100010 | 00010011011 | 11110010010

Step 5: Map the binary numbers to array indices in the dictionary
motion laugh outside latin predict mammal weird idea town mistake have kiss cream frequent great leave rocket future focus truly true fox beach venture

Step 6: Generate a 512bit seed from the words and salt
bc99a37719448b9aa4b7cdc1f3bf412bc3e17448f6bea7e3ac095d23ca63fb3c1c3e91cfcb9e77ef645de07eb619c5754ed9d4b0ef243418fdbec7db67f8b5c9
```