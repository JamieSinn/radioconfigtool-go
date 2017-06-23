# RadioConfigTool
Configuration tool for _FIRST_ Robotics Competition Robot Radios

## Technology Choice
- Using Golang provides a high speed, native binary upon compilation. This prevents decompilation to usable
source code.
- There are no pre-requisites to running a Go binary, as it has all the needed libraries and runtimes compiled
into the binary.
- Immense amount low level support and libraries. 

## How it works
This configuration tool is a re-write of the current Radio Configuration tool, with two key differences.
- No longer written in Java, so decompilation to gain access to the source is no longer possible (Go compiles to machine code.)
- Connection protocol is now encrypted - no longer possible to intercept and steal credentials without serious work.
- Embedded TFTP Server for imaging (removes the need of ap51-flash.exe)
- Listens for ARP requests to identify the model of radio

## Development Requirements

Because of the requirement of needing to intercept ARP packets, use of Google's gopacket library is required.
Gopacket requires cgo as it binds to the C `libpcap.h` header.

1. Install go_amd64 (add go binaries to your PATH)
2. Install TDM GCC x64 (add TDM-GCC binaries to your PATH) - Make sure to click the x86 and x64 option for TDM GCC. 
3. Also add TDM-GCC\x86_64-w64-mingw32\bin to your PATH
4. Install [Winpcap](https://www.winpcap.org/install/default.htm)
5. Download [Winpcap developer's pack](https://www.winpcap.org/devel.htm) and extract it to C:\
6. Find wpcap.dll and packet.dll in your PC (typically in c:\windows\system32
copy them to some other temp folder or else you'll have to supply Admin privs to the following commands
`gendef wpcap.dll`, and `gendef packet.dll`

7. Now we'll generate the static libraries files:
Run 
`dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libwpcap.a --input-def wpcap.def` 
And
`dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libpacket.a --input-def packet.def`

8. Now just copy both libwpcap.a and libpacket.a to C:\WpdPack\Lib\x64

_Win x64 compilation instructions from [Stack Overflow](https://stackoverflow.com/questions/38047858/compile-gopacket-on-windows-64bit)_


# To Do List

- ~~Create a cycle system of sorts for each team (From team number entered to configuration completed)~~
- Code cleanup, remove un-needed code, and document
- ~~Figure out a flow of operations for how the cycle will work~~
~~- Resources (images, firmwares) added to the go-bindata resource folder.~~
~~- Setup an rsrc script to inject the manifest into the final exe~~
~~- Unit testing setup~~
~~- Encryption for configuration.~~
~~- TFTP Server setup~~
~~- ARP Listener setup~~


## Software Flow (Team Use)

- Tool opens
- Team enters their number
- Instructions are on the page
- Selects either program, or image buttons.
- On selecting the program button, it sends the configuration string to the team. (pending changes to the config string and system) Return to main screen.
- On selecting the image button, listens for ARP request, get radio model, flash radio model via tftp. Return to main screen.

## Software Flow (Event Use)

- Tool opens
- Team enters their number
- Instructions are on the page
- Selects "Program"
- Listens for ARP string and gets model
- Flashes radio with image 
- Upon the radio booting up again, the radio is configured (tbd method)

## Packages

- config
  * Used to hold all configuration values that changes different parts of the program
- fileio
  * Used to read and parse OpenWRT Image files
- gui
  * Handles the GUI
- imaging
  * Used for flashing, and configuration of the router
- netconfig
  * Network based configuration, also handles network protocols.
- resources
  * Holds images, and any embedded resources
- util
  * Utility functions and methods