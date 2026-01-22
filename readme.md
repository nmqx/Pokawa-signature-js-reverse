# Pokawa Signature JS Reverse

This repository contains the reverse engineering of the Pokawa signature mechanism.

## Reversal Method

To reverse the API, I used ADB connected to a Bluestacks emulator (android 13 i believe) to intercept and analyze the traffic from the Pokawa mobile application. This allowed me to retrieve the necessary API endpoints and understand the request signing process.

[Bluestacks](https://www.bluestacks.com/)

It is important to note that the signature generation mechanism is identical between the web application and the mobile app.

## API Details

The only difference between the various Belorder APIs is a constant secret key. This secret is located in the JavaScript source code of the website or embedded within the application.

## Go Signing Server

The file `reversepokawa.go` contains a Go server that simplifies the signature generation process. It starts on port `6767` six sevenn and exposes two endpoints:

1.  **POST `/generate-canonical-string`**: Generates the canonical string required for the signature.
2.  **POST `/signature`**: Generates the final signature using the canonical string.

## Proof of Work

The file `pow.py` serves as the proof of work for this reverse engineering effort, demonstrating the valid signature generation.
