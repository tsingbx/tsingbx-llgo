package boring

// Enabled reports whether BoringCrypto is available.
// When enabled is false, all functions in this package panic.
//
// BoringCrypto is only available on linux/amd64 and linux/arm64 systems.

const available = true

const Enabled = available
