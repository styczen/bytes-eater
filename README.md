# Bytes eater

Program to consume file content and display it in hex format.

There are other well known programs like **xxd** or **hexdump**.
One problem I had with them to show the newest byte from the file,
not show multiple values buffered.
Probably there is some way to do this but I decided to write my small 
program to do this for me.

# Example usage
```
./bytes_eater -f testdata/test_file
```

Result
```
54 68 69 73 20 69 73 20 61 20 73 65 6e 74 65 6e | This.is.a.senten
63 65 2e 0a 41 6e 64 20 74 68 69 73 20 69 73 20 | ce..And.this.is.
61 20 f0 9f 92 a9 0a 48 6f 77 20 64 6f 20 79 6f | a......How.do.yo
75 20 6c 69 6b 65 20 74 68 69 73 3f 0a          | u.like.this?.
```
