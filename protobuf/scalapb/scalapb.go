package scalapb

// This file is intentionally left blank.
//
// When sharing proto definitions between Scala and Go, the Scala code generation requires the "scalapb.proto"
// import.  However, the presence of this import causes the Go compiler to expect a Go package here and generate
// an error when none is found.
//
// Since there's no sense in generating Go code for scalapb.proto, this dummy file is included instead.
//
// The only way not to have this hack would be to somehow tell the protobuf compiler that for Go files the scalapb
// dependency should be ignored.  As of the time of writing such a solution was not found.
