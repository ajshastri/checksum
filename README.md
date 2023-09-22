# Checksum
Prints checksum of the file based on Golang's crypto library and golang.org/x/crypto  

## Usage
  ### Usage: checksum [OPTION] ... [FILE]
  Prints checksum of the file based on Golang's crypto library and golang.org/x/crypto.  
  Defaults to SHA2-256 (256-bit) checksums.  

## Options:
  `-h, --help`          - Displays this help command <br />
  `-c, --checksum`      - Checksum algorithm. Defaults to SHA2 256-bit checksum.  
  `-f, --file`          - Does a checksum on the file defined. Defaults to SHA2 256-bit checksum.  

## Available checksums:
  md4 - Performs MD4 hash algorithm as defined in RFC 1320.  
  md5 - Performs MD5 hash algorithm as defined in RFC 1321.  
  sha1 - Performs SHA-1 hash algorithm as defined in RFC 3174.  
  sha2-224 / sha224 - Performs SHA2-224 hash algorithms as defined in FIPS 180-4.  
  sha2-256 / sha256 - Performs SHA2-256 hash algorithms as defined in FIPS 180-4.  
  sha2-384 / sha384 - Performs SHA2-384 hash algorithms as defined in FIPS 180-4.  
  sha2-512 / sha512 - Performs SHA2-512 hash algorithms as defined in FIPS 180-4.  
  sha3-224 - Performs SHA3-224 hash algorithms as defined in FIPS 180-4.  
  sha3-256 - Performs SHA3-256 hash algorithms as defined in FIPS 180-4.  
  sha3-384 - Performs SHA3-384 hash algorithms as defined in FIPS 180-4.  
  sha3-512 - Performs SHA3-512 hash algorithms as defined in FIPS 180-4.  
  sha2-512_224 / sha512_224 - Performs SHA2-512 hash algorithms as defined in FIPS 180-4.  
  sha2-512_256 / sha512_256 - Performs SHA2-512 hash algorithms as defined in FIPS 180-4.  
