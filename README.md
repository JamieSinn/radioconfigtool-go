# RadioConfigTool
Configuration tool for _FIRST_ Robotics Competition Robot Radios

## Technology Choice
- Using Golang provides a high speed, native binary upon compilation. This prevents decompilation to usable
source code.
- There are no pre-requisites to running a Go binary, as it has all the needed libraries and runtimes compiled
into the binary.


## Description
This tool is designed to replace the current Java based tool for the following reasons:

- The current tool is insecure as Java is compiled to bytecode, allowing for easy decompilation.
- There is no security around the connection/configuration protocol used.
  * Because of this, there are several severe security vulnerabilities that realistically cannot be fixed

This new tool allows for much greater security, removing any configuration connection between
the radio and configuration computer.

## How it works
This configuration tool works off the idea of using pre-built images for each team when at competition.
The images will be built on the FMS Server on Day 0 before teams arrive using a similar build script as
the OpenWRT-Builder project. (Grabbing all teams from the _FIRST_ API for the event)

This tool is currently designed only to replace the competition WPA Kiosk, though it could be extended to have
a mode for practice/non-event use.
