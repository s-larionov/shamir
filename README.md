## Usage
### bin/encrypt
```
cat file | bin/encrypt [args]

Usage of bin/encrypt:
  -parts int
        how many parts of the key should be created (default 5)
  -threshold int
        how many parts of the key are required for recombination (default 3)
  -output string
        filename for encoded content (default "out")
```

### bin/decrypt
```
cat file | bin/decrypt [args]

Usage of bin/decrypt:
  -output string
        filename for decoded content or empty for STDOUT
  -parts string
        comma separated parts of the key
```

## Encrypt file and share secret

Pre-requirements:
- Parts: 3 of 6
- Content: in file source.txt
- Encoded content: save to file encoded.txt

```bash
cat source.txt | bin/encrypt -parts 6 -threshhold 3 -output encoded.txt
```

## Decrypt file with required parts of the key

Pre-requirements:
- Required parts: 3
- Encoded content: in file encoded.txt
- Result should be printed to STDOUT

```bash
cat encoded.txt | bin/decrypt -parts PART1,PART2,PART3
```
