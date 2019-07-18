// Package bytes implements helpers to convert byte slices to (and from) specific data types
// using unsafe.
//
// As these conversions use the unsafe package, they _must_ be used with care in order to prevent
// undesired behavior.
//
// Usage of these functions prevents littering unsafe references all over the place.
// If you're not sure it is safe to use one of these methods: don't.
package bytes
