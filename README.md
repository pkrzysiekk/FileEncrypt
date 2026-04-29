# FileEncrypt 🔐

A simple and efficient Command Line Interface (CLI) tool for file encryption and decryption, written in **Go**.

## Features
- **AES-GCM Encryption:** Uses authenticated encryption for both confidentiality and data integrity.
- **Passphrase-based:** Securely encrypt files using your own password.
- **Cross-platform:** Compiled binary works on Windows, Linux, and macOS.

## Installation

Ensure you have [Go](https://go.dev/dl/) installed.

```bash
git clone [https://github.com/pkrzysiekk/FileEncrypt.git](https://github.com/pkrzysiekk/FileEncrypt.git)
cd FileEncrypt
go build -o fileencrypt
