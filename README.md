# CRC-32/MPEG-2 Algorithm for Golang

Golang has already crc32 package but it's include three predefined polynomials: 
* IEEE
* Castagnoli
* Koopman

If you need more information about go crc package, you can visit [this page](https://golang.org/pkg/hash/crc32/).


Even if IEEE, Castagnoli and Koopman are the most common CRC-32 polynomials, some MCUs (Microcontrollers) have CRC-32/MPEG-2 hardware implementation. You may need this algorithm if you work in an IoT project.

## Installation

`go get github.com/onurkaraoglan/go-crc32-mpeg2`

## Usage

`result := CalculateCrcMpeg2(hexData)`

Parameter must be byte slice and it returns uint32 as hex.

## Does It Work Properly?

There is a test file to make sure it works properly.

You can double check its accuracy by visiting [this page](https://crccalc.com/).

