# Planning

## Radio Side Programmer

Encryption/Decryption Process:
- Decode from base64
- Decrypt using key
- Check for illegal characters (Non AlphaNumeric) to prevent code injection (Marshall's injection exploit)
- Split based on comma delimited arguments
- UCI Commandline with args

On Connection Send:
- Model
- Image Version
- Current Configuration (encrypted and base64 encoded)
- Team?

## Event Build

Import CSV of WPA Keys - Locked to certain filepath

## Team Build

Different Encryption Key?

## Common

OpenWRT Images embedded into executable to make it harder to extract

