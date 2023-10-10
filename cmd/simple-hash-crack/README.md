# Simple sha-1 hash crack tool

Usage: ./simple-hash-crack hashes.txt guesses.txt

Each guess is hashed with sha-1 to see if it matches any of the hashes provided.

The app loads hashes to match/crack from the first input file/parameter on the command line. These hashes are loaded and saved in memory so the list of hashes should not be too large.

The guesses are loaded from the second input file/parameter on the command line. These are loaded one at a time and streamed from the input file, so the input file can be huge. We used a 100gb password list for testing: [https://github.com/ohmybahgosh/RockYou2021.txt](https://github.com/ohmybahgosh/RockYou2021.txt)

This app simply hashes each guess with sha-1, and then tries to match against any hash in the hashes.txt file.
