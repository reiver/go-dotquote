/*
Package dotquotedetect is a low level library for detecting where keys and values are in dotquote data in a []byte.

This package assumes that the data in the []byte is UTF-8.

Note that UTF-8 was designed to be a superset of ASCII. So something that is just ASCII data is automagically
also UTF-8 data.
*/
package dotquotedetect
