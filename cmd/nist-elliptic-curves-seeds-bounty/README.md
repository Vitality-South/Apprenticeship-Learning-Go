# Cracker for the NIST elliptic curves seeds bounty.

[https://words.filippo.io/dispatches/seeds-bounty/](https://words.filippo.io/dispatches/seeds-bounty/)

The app loads guesses from the guesses.txt file, transforms each guess into multiple variations of the guess (all lower case, all upper case, with/without punctuation, with/without spaces, etc.), hashes each variation with sha-1, and then tries to match against any hash in the hashes.txt file.

hashes.txt contains the NIST hashes we are trying to match against.
