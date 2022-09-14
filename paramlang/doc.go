// Package paramlang implements the parsing for the template language embedded in wac yaml
// configuration documents. It's primary feature is to handle params and param expansion into
// 'workspace' fields.
// The lexing/parsing of paramlang is handled by autogenerated code from Antlrv4. This autogenerated
// code is stored in `./parser`. No files in `./parser` should be edited by hand.
// Java is required for this go generate to work
//go:generate java -jar /usr/local/lib/antlr-4.9.2-complete.jar -Dlanguage=Go -o parser Wac.g4
//go:generate goimports -w ./parser/
package paramlang
