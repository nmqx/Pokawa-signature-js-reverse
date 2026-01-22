# Pokawa Signature JS Reverse

This repository contains the reverse engineering of the Pokawa signature mechanism.

## Reversal Method

To reverse the API, I used ADB connected to a Nox emulator to intercept and analyze the traffic from the Pokawa mobile application. This allowed me to retrieve the necessary API endpoints and understand the request signing process.

[Nox Emulator](https://www.bignox.com/)

It is important to note that the signature generation mechanism is identical between the web application and the mobile app.

## API Details

The only difference between the various Belorder APIs is a constant secret key. This secret is located in the JavaScript source code of the website or embedded within the application.

## Proof of Work

The file `pow.py` serves as the proof of work for this reverse engineering effort, demonstrating the valid signature generation.
