package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
	"hash"
	"io"
	"log"
	"os"
	"strings"
)

const checksumVersion = "0.3.0"

var (
	sumPtr   string
	fileName string
	h        hash.Hash
)

func version() {
	fmt.Println("Checksum version:", checksumVersion)
	os.Exit(0)
}

func help(exitCode int) {
	fmt.Println("Usage: checksum [OPTION] ... [FILE]")
	fmt.Printf("  Prints checksum of the file based on Golang's crypto library and golang.org/x/crypto.\n" +
		"  Defaults to SHA2-256 (256-bit) checksums.\n\n")
	fmt.Println("Options:  ")
	fmt.Println("  -h, --help          Displays this help command")
	fmt.Println("  -V, --version       Displays the current version")
	fmt.Println("  -s, --checksum      Checksum algorithm. Defaults to SHA2 256-bit checksum.")
	fmt.Printf("  -f, --file          Does a checksum on the file defined. Defaults to SHA2 256-bit checksum.\n\n")
	fmt.Println("Available checksums:")
	fmt.Println("  md4 - Performs MD4 hash algorithm as defined in RFC 1320.")
	fmt.Println("  md5 - Performs MD5 hash algorithm as defined in RFC 1321.")
	fmt.Println("  sha1 - Performs SHA-1 hash algorithm as defined in RFC 3174.")
	fmt.Println("  sha224 / sha2-224 - Performs SHA2-224 hash algorithms as defined in FIPS 180-4.")
	fmt.Println("  sha256 / sha2-256 - Performs SHA2-256 hash algorithms as defined in FIPS 180-4.")
	fmt.Println("  sha384 / sha2-384 - Performs SHA2-384 hash algorithms as defined in FIPS 180-4.")
	fmt.Println("  sha512 / sha2-512 - Performs SHA2-512 hash algorithms as defined in FIPS 180-4.")
	fmt.Println("  sha3-224 - Performs SHA3-224 hash algorithms as defined in FIPS 180-4.")
	fmt.Println("  sha3-256 - Performs SHA3-256 hash algorithms as defined in FIPS 180-4.")
	fmt.Println("  sha512_224 / sha2-512_224 - Performs SHA2-512 hash algorithms as defined in FIPS 180-4.")
	fmt.Println("  sha512_256 / sha2-512_256 - Performs SHA2-512 hash algorithms as defined in FIPS 180-4.")

	os.Exit(exitCode)
}

func checksumCalc(sumPtr string, fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		//log.Fatal(err)
		if os.IsNotExist(err) {
			log.Fatal("File does not exist, please refer to the help.\n\n")
			//help(1)
		}
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			//log.Fatal(err)
		}
	}(f)

	// h := sha256.New() // Defaults to SHA256Sum
	switch sumPtr {
	case "md4":
		// fmt.Println("Switch case", sumPtr)
		h = md4.New()
	case "md5":
		// fmt.Println("Switch case", sumPtr)
		h = md5.New()
	case "sha1":
		// fmt.Println("Switch case", sumPtr)
		h = sha1.New()
	case "sha2-224", "sha224":
		// fmt.Println("Switch case", sumPtr)
		h = sha256.New224()
	case "sha256", "sha2-256":
		// fmt.Println("Switch case", sumPtr)
		h = sha256.New()
	case "sha384", "sha2-384":
		// fmt.Println("Switch case", sumPtr)
		h = sha512.New384()
	case "sha512", "sha2-512":
		// fmt.Println("Switch case", sumPtr)
		h = sha512.New()
	case "sha2-512_224", "sha512_224":
		// fmt.Println("Switch case", sumPtr)
		h = sha512.New512_224()
	case "sha2-512_256", "sha512_256":
		// fmt.Println("Switch case", sumPtr)
		h = sha512.New512_256()
	case "sha3-224":
		// fmt.Println("Switch case", sumPtr)
		h = sha3.New224()
	case "sha3-256":
		// fmt.Println("Switch case", sumPtr)
		h = sha3.New256()
	case "sha3-384":
		// fmt.Println("Switch case", sumPtr)
		h = sha3.New384()
	case "sha3-512":
		// fmt.Println("Switch case", sumPtr)
		h = sha3.New512()
	default:
		// fmt.Println("Unidentified Checksum. Please refer to the help document:")
		help(2)
	}
	if _, err := io.Copy(h, f); err != nil {
		// log.Fatal(err)
		help(1)
	}
	fmt.Printf("%x  %s\n", h.Sum(nil), fileName)
}

func main() {

	flag.Usage = func() { help(2) }
	verVer := flag.Bool("V", false, "A bool flag")
	verVer2 := flag.Bool("version", false, "A bool flag")
	flag.StringVar(&sumPtr, "checksum", "sha2-256", "Checksum to be calculated")
	flag.StringVar(&sumPtr, "s", "sha2-256", "Checksum to be calculated")
	flag.StringVar(&fileName, "f", "", "")
	flag.StringVar(&fileName, "file", "", "")
	flag.Parse()

	if *verVer || *verVer2 {
		version()
	}

	checksumCalc(strings.ToLower(sumPtr), fileName)
}
